package secret

import (
	"context"
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	datapb "github.com/ksusonic/gophkeeper/proto/data"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Storage interface for managins secrets in storage
type Storage interface {
	SetSecret(ctx context.Context, secret *models.Secret) error
	GetSecret(ctx context.Context, userID, name string) (*models.Secret, error)
	GetAllSecrets(ctx context.Context, userID string) ([]models.Secret, error)
	UpdateSecret(ctx context.Context, secret *models.Secret) error
	RemoveSecret(ctx context.Context, userID, name string) (bool, error)
}

// Controller for secrets
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

// SetSecret creates secret or updates curretnly existing
func (c *Controller) SetSecret(ctx context.Context, claims *models.UserClaims, secret *datapb.Secret) (*servicepb.SetSecretResponse, error) {
	existingSecret, err := c.storage.GetSecret(ctx, claims.UserID, secret.GetName())
	if err != nil {
		c.logger.Error("error getting secret: %v", err)
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
	if existingSecret == nil {
		// create new
		err := c.storage.SetSecret(ctx, &models.Secret{
			UserID: claims.UserID,
			Name:   secret.GetName(),
			Meta:   secret.GetMeta().AsMap(),
			Data:   encryptedData,
		})
		if err != nil {
			c.logger.Error("error saving secret: %v", err)
			return nil, status.Errorf(codes.Internal, "could not save secret: %v", err)
		}
		return &servicepb.SetSecretResponse{}, nil
	} else {
		// update existing
		existingSecret.Version++
		// merge meta
		for k, v := range secret.Meta.AsMap() {
			existingSecret.Meta[k] = v
		}
		existingSecret.Data = encryptedData
		err := c.storage.UpdateSecret(ctx, existingSecret)
		if err != nil {
			c.logger.Error("could not update existing secret: %v", err)
			return nil, status.Errorf(codes.Internal, "could not update secret: %v", err)
		}
		return &servicepb.SetSecretResponse{}, nil
	}
}

// GetSecret gets secret from storage or returns error NotFound
func (c *Controller) GetSecret(ctx context.Context, claims *models.UserClaims, name string) (*servicepb.GetSecretResponse, error) {
	secret, err := c.storage.GetSecret(ctx, claims.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get secret: %v", err)
	}
	if secret == nil {
		return nil, status.Errorf(codes.NotFound, "secret %s not found in users secrets", name)
	}

	decryptedSecret, err := crypta.DecryptSecret(c.crypta, secret)
	if err != nil {
		c.logger.Error("could not decrypt secret: %v", err)
		return nil, status.Errorf(codes.Internal, "unexpected error")
	}

	return &servicepb.GetSecretResponse{
		Secret: decryptedSecret,
	}, nil
}

// GetAllSecrets returns all secrets of user
func (c *Controller) GetAllSecrets(ctx context.Context, claims *models.UserClaims) (*servicepb.GetAllSecretsResponse, error) {
	secrets, err := c.storage.GetAllSecrets(ctx, claims.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get secrets: %v", err)
	}
	var protoSecrets = make([]*datapb.Secret, len(secrets))
	eg := errgroup.Group{}
	for i := range secrets {
		f := func(i int) func() (err error) {
			return func() (err error) {
				protoSecrets[i], err = crypta.DecryptSecret(c.crypta, &secrets[i])
				return err
			}
		}
		eg.Go(f(i))
	}
	if err := eg.Wait(); err != nil {
		c.logger.Error("could not decrypt secret: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &servicepb.GetAllSecretsResponse{
		Secrets: protoSecrets,
	}, nil
}

// RemoveSecret removes secret by name
func (c *Controller) RemoveSecret(ctx context.Context, claims *models.UserClaims, name string) (*servicepb.RemoveSecretResponse, error) {
	found, err := c.storage.RemoveSecret(ctx, claims.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get secret: %v", err)
	}
	if !found {
		return nil, status.Errorf(codes.NotFound, "secret %s not found and not deleted", name)
	}

	return &servicepb.RemoveSecretResponse{}, nil
}
