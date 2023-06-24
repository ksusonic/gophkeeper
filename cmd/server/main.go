package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/db"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/server"

	"github.com/mborders/logmatic"
)

const shutdownTimeout = 5 * time.Second

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	logger := getLogger(cfg.DebugMode)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := runServer(ctx, cfg, logger); err != nil {
		logger.Fatal("server error: %v", err)
	}
}

func getLogger(debug bool) logging.Logger {
	l := logmatic.NewLogger()
	if debug {
		l.SetLevel(logmatic.DEBUG)
	}
	return l
}

func runServer(ctx context.Context, cfg *config.Config, logger logging.Logger) error {
	/*storage*/
	_, err := db.NewDB(cfg.Server.DatabaseDsn)
	if err != nil {
		logger.Fatal("db error: %v", err)
	}
	srv := server.NewGrpcServer(&cfg.Server, logger)

	go srv.ListenAndServe(cfg.Server.Address)

	<-ctx.Done()
	logger.Info("caught stop signal")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	grpcShutdown := make(chan struct{}, 1)
	go func() {
		srv.GracefulStop()
		close(grpcShutdown)
	}()

	select {
	case <-shutdownCtx.Done():
		return fmt.Errorf("server shutdown: %w", ctx.Err())
	case <-grpcShutdown:
		logger.Info("grpc server down")
	}

	return nil
}
