package db

import (
	"context"
	"errors"
	"time"

	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultTimeout = time.Second * 3

type DB struct {
	orm    *gorm.DB
	logger logging.Logger
}

func NewDB(dsn string, logger logging.Logger) (*DB, error) {
	orm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Error connecting to the database...", err)
	}

	if err := orm.AutoMigrate(models.AllModels...); err != nil {
		logger.Fatal("Error in automigrations: %v", err)
	}
	logger.Info("db automigration ok")

	return &DB{
		orm: orm,
	}, nil
}

func (d *DB) SaveUser(ctx context.Context, user *models.User) error {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	return d.orm.WithContext(tCtx).Save(user).Error
}

func (d *DB) GetUser(ctx context.Context, username string) (*models.User, error) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	user := &models.User{}
	tx := d.orm.WithContext(tCtx).Where("username = ?", username).Take(user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return user, nil
}

func (d *DB) SetSecret(ctx context.Context, secret *models.Secret) error {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	tx := d.orm.WithContext(tCtx).Create(secret)
	return tx.Error
}

func (d *DB) GetSecret(ctx context.Context, userID string, name string) (*models.Secret, error) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	secret := &models.Secret{}
	tx := d.orm.WithContext(tCtx).Where("user_id = ? AND name = ?", userID, name).Take(secret)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}
	return secret, nil
}

func (d *DB) GetAllSecrets(ctx context.Context, userID string) ([]models.Secret, error) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	var secrets []models.Secret
	tx := d.orm.WithContext(tCtx).Where("user_id = ?", userID).Find(&secrets)
	return secrets, tx.Error
}

func (d *DB) UpdateSecret(ctx context.Context, secret *models.Secret) error {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	return d.orm.WithContext(tCtx).Save(secret).Error
}

func (d *DB) RemoveSecret(ctx context.Context, userID, name string) (bool, error) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	secret := models.Secret{}
	tx := d.orm.WithContext(tCtx).Where("user_id = ? AND name = ?", userID, name).Delete(&secret)
	return tx.RowsAffected > 0, tx.Error
}
