package grpc

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/controller/auth"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"

	"google.golang.org/grpc"
)

type AuthControllerGrpc struct {
	controller *auth.Controller
	servicepb.UnimplementedAuthServiceServer
}

func NewAuthControllerGrpc(userStorage auth.UserStorage, jwtManager *crypta.JWTManager, logger logging.Logger) *AuthControllerGrpc {
	return &AuthControllerGrpc{
		controller: auth.NewController(userStorage, jwtManager, logger),
	}
}

func (a *AuthControllerGrpc) RegisterService(srv *grpc.Server) {
	servicepb.RegisterAuthServiceServer(srv, a)
}

func (a *AuthControllerGrpc) ServiceName() string {
	return servicepb.AuthService_ServiceDesc.ServiceName
}

func (a *AuthControllerGrpc) Register(ctx context.Context, request *servicepb.RegisterRequest) (*servicepb.RegisterResponse, error) {
	return a.controller.Register(ctx, request.GetLogin(), request.GetPassword())
}

func (a *AuthControllerGrpc) Login(ctx context.Context, request *servicepb.LoginRequest) (*servicepb.LoginResponse, error) {
	return a.controller.Login(ctx, request.GetLogin(), request.GetPassword())
}
