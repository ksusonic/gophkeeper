package server

import (
	"net"
	"runtime/debug"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/logging"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcController interface {
	Name() string
	RegisterService(srv *grpc.Server)
}

type GrpcServer struct {
	srv    *grpc.Server
	logger logging.Logger
	cfg    *config.ServerConfig
}

func NewGrpcServer(cfg *config.ServerConfig, logger logging.Logger, authFunc auth.AuthFunc, controllers ...GrpcController) *GrpcServer {
	server := &GrpcServer{
		srv: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				auth.UnaryServerInterceptor(authFunc),
				recovery.UnaryServerInterceptor(
					recovery.WithRecoveryHandler(panicRecoveryHandler(logger)),
				),
			),
			grpc.ChainStreamInterceptor(
				auth.StreamServerInterceptor(authFunc),
			)),
		cfg:    cfg,
		logger: logger,
	}
	for _, controller := range controllers {
		logger.Info("Registered %s controller", controller.Name())
		controller.RegisterService(server.srv)
	}
	return server
}

func (s *GrpcServer) ListenAndServe(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		s.logger.Fatal("Cannot listen: %v", err)
	}

	s.logger.Info("GRPC server Listening at %v", lis.Addr())
	if err := s.srv.Serve(lis); err != nil {
		s.logger.Fatal("Failed to serve: %v", err)
	}
}

func (s *GrpcServer) GracefulStop() {
	s.srv.GracefulStop()
}

func panicRecoveryHandler(logger logging.Logger) func(p any) error {
	return func(p any) error {
		errorID := uuid.New()
		logger.Error("errorID: %s got panic: %v, stack: %s", errorID, p, debug.Stack())
		return status.Errorf(codes.Internal, "internal error id: %s", errorID)
	}
}
