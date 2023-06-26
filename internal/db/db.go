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

func (d *DB) SaveUser(ctx context.Context, user *models.User) models.StorageQueryResult {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	return d.orm.WithContext(tCtx).Save(user).Error
}

func (d *DB) GetUser(ctx context.Context, username string) (*models.User, models.StorageQueryResult) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	user := &models.User{}
	tx := d.orm.WithContext(tCtx).Where("username = ?", username).Take(user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, models.ErrorNotExists
		}
		return nil, tx.Error
	}
	return user, nil
}

func (d *DB) SetSecret(ctx context.Context, secret *models.Secret) error {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	tx := d.orm.WithContext(tCtx).Save(secret)
	return tx.Error
}

func (d *DB) UserHasSecret(ctx context.Context, userID string, name string) (bool, error) {
	tCtx, timeout := context.WithTimeout(ctx, defaultTimeout)
	defer timeout()

	var count int64
	tx := d.orm.WithContext(tCtx).Model(models.Secret{}).Where("user_id = ? AND name = ?", userID, name).Count(&count)
	if tx.Error != nil {
		return false, tx.Error
	}
	return count > 0, nil
}

func (d *DB) GetSecret(ctx context.Context, userID string, name string) (*models.Secret, error) {
	//TODO implement me
	panic("implement me")
}
