package models

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   uint
	Username string `json:"username"`
}
