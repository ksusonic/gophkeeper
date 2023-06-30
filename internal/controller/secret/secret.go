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
	"google.golang.org/protobuf/types/known/structpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Storage interface {
	SetSecret(ctx context.Context, secret *models.Secret) error
	GetSecret(ctx context.Context, userID, name string) (*models.Secret, error)
	UpdateSecret(ctx context.Context, secret *models.Secret) error
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
			Name:   secret.Name,
			Meta:   secret.Meta.AsMap(),
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

func (c *Controller) GetSecret(ctx context.Context, claims *models.UserClaims, name string) (*servicepb.GetSecretResponse, error) {
	secret, err := c.storage.GetSecret(ctx, claims.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get secret: %v", err)
	}
	if secret == nil {
		return nil, status.Errorf(codes.NotFound, "secret %s not found in users secrets", name)
	}

	decryptedData, err := c.crypta.Decrypt(secret.Data)
	if err != nil {
		c.logger.Error("could not decrypt %s secret secretData: %v", name, err)
		return nil, status.Error(codes.Internal, "could not decrypt secretData. sorry :-(")
	}
	secretData := &datapb.Secret_Data{}
	if err = proto.Unmarshal(decryptedData, secretData); err != nil {
		c.logger.Error("could not unmarshall %s secretData: %v", name, err)
		return nil, status.Error(codes.Internal, "unexpected unmarshalling error")
	}
	meta, err := structpb.NewStruct(secret.Meta)
	if err != nil {
		c.logger.Error("could not create struct from map: %v", err)
		return nil, status.Error(codes.Internal, "unexpected error")
	}

	return &servicepb.GetSecretResponse{
		Secret: &datapb.Secret{
			Name:       secret.Name,
			Meta:       meta,
			SecretData: secretData,
		},
	}, nil
}

func (c *Controller) GetAllSecrets(ctx context.Context, claims *models.UserClaims) (*servicepb.GetAllSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}
