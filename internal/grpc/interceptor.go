package grpc

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ksusonic/gophkeeper/internal/crypta"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager         *crypta.JWTManager
	ignoreServiceNames []string
}

func NewAuthInterceptor(jwtManager *crypta.JWTManager, ignoreServiceNames ...string) *Interceptor {
	return &Interceptor{
		jwtManager:         jwtManager,
		ignoreServiceNames: ignoreServiceNames,
	}
}

func (i *Interceptor) AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	claims, err := i.jwtManager.Verify(token)
	if err != nil {
		if errors.Is(errors.Unwrap(err), jwt.ErrTokenInvalidClaims) {
			return ctx, status.Error(codes.Unauthenticated, "access token is expired")
		}
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return context.WithValue(ctx, ClaimsKey, claims), nil
}

func (i *Interceptor) Match(_ context.Context, fullMethod interceptors.CallMeta) bool {
	for _, servieName := range i.ignoreServiceNames {
		if fullMethod.Service == servieName {
			return false
		}
	}
	return true
}
