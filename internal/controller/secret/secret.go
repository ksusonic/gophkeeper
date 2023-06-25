package secret

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	"github.com/ksusonic/gophkeeper/proto/data"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
)

type Storage interface {
	SetSecret(ctx context.Context, secret *models.Secret) error
	GetSecret(ctx context.Context, userID, name string)
	HasSecret(ctx context.Context)
}

type Controller struct {
	logger  logging.Logger
	storage Storage
}

func NewController(storage Storage, logger logging.Logger) *Controller {
	return &Controller{
		storage: storage,
		logger:  logger,
	}
}

func (c *Controller) SetSecret(ctx context.Context, data *data.Secret) (*servicepb.SetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetSecret(ctx context.Context, name string) (*servicepb.GetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetAllSecrets(ctx context.Context) (*servicepb.GetAllSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}
