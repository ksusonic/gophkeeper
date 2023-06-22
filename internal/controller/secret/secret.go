package secret

import (
	"context"

	"github.com/ksusonic/gophkeeper/proto/data"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
)

type SecretsController struct {
}

func (c *SecretsController) SetSecret(ctx context.Context, data *data.Secret) (*servicepb.SetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *SecretsController) GetSecret(ctx context.Context, name string) (*servicepb.GetSecretResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *SecretsController) GetAllSecrets(ctx context.Context) (*servicepb.GetAllSecretsResponse, error) {
	//TODO implement me
	panic("implement me")
}
