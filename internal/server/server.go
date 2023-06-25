package server

import (
	"net"
	"runtime/debug"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/logging"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcController interface {
	ServiceName() string
	RegisterService(srv *grpc.Server)
}

type GrpcServer struct {
	srv    *grpc.Server
	logger logging.Logger
	cfg    *config.ServerConfig
}

func NewGrpcServer(
	cfg *config.ServerConfig,
	logger logging.Logger,
	authFunc auth.AuthFunc,
	needAuthFilter selector.MatchFunc,
) *GrpcServer {
	return &GrpcServer{
		srv: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFunc), needAuthFilter),
				recovery.UnaryServerInterceptor(
					recovery.WithRecoveryHandler(panicRecoveryHandler(logger)),
				),
			),
			grpc.ChainStreamInterceptor(
				selector.StreamServerInterceptor(auth.StreamServerInterceptor(authFunc), needAuthFilter),
				recovery.StreamServerInterceptor(
					recovery.WithRecoveryHandler(panicRecoveryHandler(logger)),
				),
			)),
		cfg:    cfg,
		logger: logger,
	}
}

func (s *GrpcServer) RegisterController(controller GrpcController) {
	s.logger.Info("Registered %s", controller.ServiceName())
	controller.RegisterService(s.srv)
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
