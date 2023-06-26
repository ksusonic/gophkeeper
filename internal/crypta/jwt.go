package crypta

import (
	"fmt"
	"time"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

const (
	Issuer = "gophkeeper"
)

type JWTManager struct {
	saltKey  string
	tokenTTL time.Duration
}

func NewJWTManager(config config.AuthConfig) *JWTManager {
	return &JWTManager{
		saltKey:  config.SaltKey,
		tokenTTL: config.TokenTTL,
	}
}

func (manager *JWTManager) Generate(user *models.User) (string, error) {
	claims := models.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenTTL)),
			Issuer:    Issuer,
		},
		UserID:   user.ID,
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.saltKey))
}

func (manager *JWTManager) Verify(accessToken string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&models.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.saltKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
