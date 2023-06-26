package grpc

import (
	"context"
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ClaimsKey struct{}

func retrieveClaims(ctx context.Context) (*models.UserClaims, error) {
	if anyToken := ctx.Value(ClaimsKey); anyToken != nil {
		token, ok := anyToken.(*models.UserClaims)
		if !ok {
			return nil, status.Errorf(codes.Internal, "could not retrieve token from context: %v", anyToken)
		}
		return token, nil
	}
	return nil, fmt.Errorf("token not found")
}
