package auth

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/models"
)

// UserStorage interface for saving user data
type UserStorage interface {
	SaveUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string) (*models.User, error)
}
