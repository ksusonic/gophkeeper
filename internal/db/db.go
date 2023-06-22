package db

import (
	"log"

	"github.com/ksusonic/gophkeeper/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	if err := db.AutoMigrate(models.AllModels...); err != nil {
		log.Fatalf("Error in automigrations: %v", err)
	}

	return &DB{
		DB: db,
	}, nil
}
