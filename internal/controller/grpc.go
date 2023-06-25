package controller

import (
	"context"

	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/ksusonic/gophkeeper/internal/controller/auth"
	"github.com/ksusonic/gophkeeper/internal/controller/secret"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"
)

var whitelistedAuthPaths = map[string]bool{
	"/service.AuthService/": true,
}

type AuthControllerGrpc struct {
	controller  *auth.Controller
	interceptor *Interceptor
	servicepb.UnimplementedAuthServiceServer
}

func NewAuthControllerGrpc(userStorage auth.UserStorage, jwtManager *crypta.JWTManager, logger logging.Logger) *AuthControllerGrpc {
	return &AuthControllerGrpc{
		controller:  auth.NewController(userStorage, jwtManager, logger),
		interceptor: NewAuthInterceptor(jwtManager, whitelistedAuthPaths),
	}
}

func (a *AuthControllerGrpc) RegisterService(srv *grpc.Server) {
	servicepb.RegisterAuthServiceServer(srv, a)
}

func (a *AuthControllerGrpc) Name() string {
	return "Authentication grpc-controller"
}

func (a *AuthControllerGrpc) AuthFunc() grpcAuth.AuthFunc {
	return a.interceptor.AuthFunc
}

func (a *AuthControllerGrpc) Register(ctx context.Context, request *servicepb.RegisterRequest) (*servicepb.RegisterResponse, error) {
	return a.controller.Register(ctx, request.GetLogin(), request.GetPassword())
}

func (a *AuthControllerGrpc) Login(ctx context.Context, request *servicepb.LoginRequest) (*servicepb.LoginResponse, error) {
	return a.controller.Login(ctx, request.GetLogin(), request.GetPassword())
}

type SecretControllerGrpc struct {
	controller *secret.Controller
	servicepb.UnimplementedSecretServiceServer
}

func NewSecretControllerGrpc(userStorage secret.Storage, logger logging.Logger) *SecretControllerGrpc {
	return &SecretControllerGrpc{
		controller: secret.NewController(userStorage, logger),
	}
}

func (s *SecretControllerGrpc) RegisterService(srv *grpc.Server) {
	servicepb.RegisterSecretServiceServer(srv, s)
}

func (s *SecretControllerGrpc) Name() string {
	return "Secrets grpc-controller"
}

func (s *SecretControllerGrpc) SetSecret(ctx context.Context, request *servicepb.SetSecretRequest) (*servicepb.SetSecretResponse, error) {
	return s.controller.SetSecret(ctx, request.GetData())
}

func (s *SecretControllerGrpc) GetSecret(ctx context.Context, request *servicepb.GetSecretRequest) (*servicepb.GetSecretResponse, error) {
	return s.controller.GetSecret(ctx, request.GetName())
}

func (s *SecretControllerGrpc) GetAllSecrets(ctx context.Context, _ *emptypb.Empty) (*servicepb.GetAllSecretsResponse, error) {
	return s.controller.GetAllSecrets(ctx)
}
