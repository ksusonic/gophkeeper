package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager           *crypta.JWTManager
	whiteListedAuthPaths map[string]bool
}

func NewAuthInterceptor(jwtManager *crypta.JWTManager, whiteListedAuthPaths map[string]bool) *Interceptor {
	return &Interceptor{
		jwtManager,
		whiteListedAuthPaths,
	}
}

func (i *Interceptor) AuthFunc(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if !ok {
		return ctx, fmt.Errorf("can't get method from request")
	}
	if i.isWhitelisted(method) {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	_, err = i.jwtManager.Verify(token)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return ctx, nil
}

func (i *Interceptor) isWhitelisted(method string) bool {
	for path := range i.whiteListedAuthPaths {
		if strings.HasPrefix(method, path) {
			return true
		}
	}
	return false
}
