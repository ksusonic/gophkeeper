package server

import (
	"net"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"google.golang.org/grpc"
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

func NewGrpcServer(cfg *config.ServerConfig, logger logging.Logger, controllers ...GrpcController) *GrpcServer {
	server := &GrpcServer{
		srv:    grpc.NewServer(),
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

	if err := s.srv.Serve(lis); err != nil {
		s.logger.Fatal("Failed to serve: %v", err)
	}
}

func (s *GrpcServer) GracefulStop() {
	s.srv.GracefulStop()
}
