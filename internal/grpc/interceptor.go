package grpc

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	jwtManager *crypta.JWTManager
	filterFunc selector.MatchFunc
}

func NewAuthInterceptor(jwtManager *crypta.JWTManager, ignoreServiceNames ...string) *Interceptor {
	return &Interceptor{
		jwtManager: jwtManager,
		filterFunc: func(ctx context.Context, fullMethod string) bool {
			fullMethod = strings.TrimPrefix(fullMethod, "/")
			for _, servieName := range ignoreServiceNames {
				if strings.HasPrefix(fullMethod, servieName) {
					return false
				}
			}
			return true
		},
	}
}

func (i *Interceptor) AuthFunc(ctx context.Context) (context.Context, error) {
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

func (i *Interceptor) FilterFunc(ctx context.Context, fullMethod string) bool {
	return i.filterFunc(ctx, fullMethod)
}
