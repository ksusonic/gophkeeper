package grpc

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/controller/secret"
	"github.com/ksusonic/gophkeeper/internal/logging"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SecretControllerGrpc struct {
	controller *secret.Controller
	servicepb.UnimplementedSecretServiceServer
}

func (s *SecretControllerGrpc) ServiceName() string {
	return servicepb.AuthService_ServiceDesc.ServiceName
}

func NewSecretControllerGrpc(userStorage secret.Storage, logger logging.Logger) *SecretControllerGrpc {
	return &SecretControllerGrpc{
		controller: secret.NewController(userStorage, logger),
	}
}

func (s *SecretControllerGrpc) RegisterService(srv *grpc.Server) {
	servicepb.RegisterSecretServiceServer(srv, s)
}

func (s *SecretControllerGrpc) SetSecret(ctx context.Context, request *servicepb.SetSecretRequest) (*servicepb.SetSecretResponse, error) {
	return s.controller.SetSecret(ctx, request.GetSecret())
}

func (s *SecretControllerGrpc) GetSecret(ctx context.Context, request *servicepb.GetSecretRequest) (*servicepb.GetSecretResponse, error) {
	return s.controller.GetSecret(ctx, request.GetName())
}

func (s *SecretControllerGrpc) GetAllSecrets(ctx context.Context, _ *emptypb.Empty) (*servicepb.GetAllSecretsResponse, error) {
	return s.controller.GetAllSecrets(ctx)
}
