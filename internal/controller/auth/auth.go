package auth

import (
	"context"

	servicepb "github.com/ksusonic/gophkeeper/proto/service"
)

type Controller struct {
	userStorage UserStorage
}

func (c *Controller) Register(ctx context.Context, email, password string) (*servicepb.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Controller) Login(ctx context.Context, email string, password string) (*servicepb.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
