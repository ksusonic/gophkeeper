package auth

import "github.com/ksusonic/gophkeeper/internal/models"

type UserStorage interface {
	SaveUser() error
	GetUser() (models.User, error)
}
