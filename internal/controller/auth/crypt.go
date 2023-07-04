package auth

import (
	"github.com/ksusonic/gophkeeper/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword - calculates hash and error if not success
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// IsCorrectPassword returns true if password hashes are equal
func IsCorrectPassword(user *models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
