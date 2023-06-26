package secret

import (
	"context"
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	"github.com/ksusonic/gophkeeper/proto/data"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Storage interface {
	SetSecret(ctx context.Context, secret *models.Secret) error
	GetSecret(ctx context.Context, userID, name string) (*models.Secret, error)
	UserHasSecret(ctx context.Context, userID, name string) (bool, error)
}

type Controller struct {
	logger  logging.Logger
	storage Storage
	crypta  *crypta.Service
}

func NewController(cfg config.SecretsConfig, storage Storage, logger logging.Logger) (*Controller, error) {
	cryptaService, err := crypta.NewService(cfg.KeysDir)
	if err != nil {
		return nil, fmt.Errorf("could not create secrets controller: %w", err)
	}

	return &Controller{
		storage: storage,
		logger:  logger,
		crypta:  cryptaService,
	}, nil
}

func (c *Controller) SetSecret(ctx context.Context, claims *models.UserClaims, secret *data.Secret) (*servicepb.SetSecretResponse, error) {
	ok, err := c.storage.UserHasSecret(ctx, claims.UserID, secret.GetName())
	if err != nil {
		c.logger.Error("could not check hasUser: %v", err)
		return nil, status.Error(codes.Internal, "unexpected storage error")
	}

	bytesData, err := proto.Marshal(secret.GetSecretData())
	if err != nil {
		c.logger.Error("could not check marshall proto secretData: %v", err)
		return nil, status.Error(codes.Internal, "unexpected server error")
	}
	encryptedData, err := c.crypta.Encrypt(bytesData)
	if err != nil {
		c.logger.Error("could not check encrypt proto secretData: %v", err)
		return nil, status.Error(codes.Internal, "unexpected server error")
	}
	if !ok {
		// create new
		err := c.storage.SetSecret(ctx, &models.Secret{
			UserID: claims.UserID,
			Name:   secret.Name,
			Meta:   secret.Meta.AsMap(),
			Data:   encryptedData,
		})
		if err != nil {
			c.logger.Error("error saving secret: %v", err)
			return nil, fmt.Errorf("could not save secret: %w", err)
		}
		return &servicepb.SetSecretResponse{}, nil
	} else {
		// update exists
		return &servicepb.SetSecretResponse{}, nil
	}
}

func (c *Controller) GetSecret(ctx context.Context, claims *models.UserClaims, name string) (*servicepb.GetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) GetAllSecrets(ctx context.Context, claims *models.UserClaims) (*servicepb.GetAllSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}
