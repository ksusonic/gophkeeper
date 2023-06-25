package auth

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/models"
)

type UserStorage interface {
	SaveUser(ctx context.Context, user *models.User) models.StorageQueryResult
	GetUser(ctx context.Context, username string) (*models.User, models.StorageQueryResult)
}

type UserNotFound error
