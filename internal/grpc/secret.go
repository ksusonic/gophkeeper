package grpc

import (
	"context"
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/controller/secret"
	"github.com/ksusonic/gophkeeper/internal/logging"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"

	"google.golang.org/grpc"
)

type SecretControllerGrpc struct {
	controller *secret.Controller
	servicepb.UnimplementedSecretServiceServer
}

func NewSecretControllerGrpc(cfg config.SecretsConfig, userStorage secret.Storage, logger logging.Logger) (*SecretControllerGrpc, error) {
	controller, err := secret.NewController(cfg, userStorage, logger)
	if err != nil {
		return nil, fmt.Errorf("error creating SecretController: %w", err)
	}
	return &SecretControllerGrpc{controller: controller}, nil
}

func (s *SecretControllerGrpc) ServiceName() string {
	return servicepb.AuthService_ServiceDesc.ServiceName
}

func (s *SecretControllerGrpc) RegisterService(srv *grpc.Server) {
	servicepb.RegisterSecretServiceServer(srv, s)
}

func (s *SecretControllerGrpc) SetSecret(ctx context.Context, request *servicepb.SetSecretRequest) (*servicepb.SetSecretResponse, error) {
	claims, err := retrieveClaims(ctx)
	if err != nil {
		return nil, err
	}
	return s.controller.SetSecret(ctx, claims, request.GetSecret())
}

func (s *SecretControllerGrpc) GetSecret(ctx context.Context, request *servicepb.GetSecretRequest) (*servicepb.GetSecretResponse, error) {
	claims, err := retrieveClaims(ctx)
	if err != nil {
		return nil, err
	}
	return s.controller.GetSecret(ctx, claims, request.GetName())
}

func (s *SecretControllerGrpc) GetAllSecrets(ctx context.Context, _ *servicepb.GetAllSecretsRequest) (*servicepb.GetAllSecretsResponse, error) {
	claims, err := retrieveClaims(ctx)
	if err != nil {
		return nil, err
	}
	return s.controller.GetAllSecrets(ctx, claims)
}
