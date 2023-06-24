package models

import (
	"gorm.io/gorm"
)

var AllModels = []interface{}{
	User{},
	Secret{},
}

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique"`
	PasswordHash string `gorm:"not null"`

	Secrets []Secret
}

type Secret struct {
	gorm.Model
	UserID  uint `gorm:"primarykey"`
	Version int32
	Meta    []byte
	Data    []byte
}
