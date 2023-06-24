package auth

import (
	"github.com/ksusonic/gophkeeper/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func IsCorrectPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
