package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var AllModels = []interface{}{
	User{},
	Secret{},
}

type User struct {
	gorm.Model
	ID           string `gorm:"type:uuid;default:gen_random_uuid()"`
	Username     string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`

	Secrets []Secret
}

type Secret struct {
	gorm.Model
	Name    string `gorm:"index"`
	Version int32  `gorm:"default:1"`
	Meta    datatypes.JSONMap
	Data    []byte

	UserID string `gorm:"type:uuid"`
}
