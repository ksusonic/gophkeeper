package auth

import (
	"context"
	"errors"

	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	userStorage UserStorage
	jwtManager  *crypta.JWTManager
	logger      logging.Logger
}

func NewController(userStorage UserStorage, jwtManager *crypta.JWTManager, logger logging.Logger) *Controller {
	return &Controller{
		userStorage: userStorage,
		jwtManager:  jwtManager,
		logger:      logger,
	}
}

func (c *Controller) Register(ctx context.Context, username, password string) (*servicepb.RegisterResponse, error) {
	if user, err := c.userStorage.GetUser(ctx, username); err != nil && !errors.Is(err, models.ErrorNotExists) {
		c.logger.Error("unexpected error from storage: %v", err)
		return nil, status.Error(codes.Internal, "unexpected error storage processing")
	} else if user != nil {
		return nil, status.Errorf(codes.AlreadyExists, "username %s already taken", username)
	}

	hashedPassword, err := HashPassword(password)
	if err != nil {
		c.logger.Error("Can't hash password: %v", err)
		return nil, status.Error(codes.Internal, "unexpected error in password processing")
	}
	user := &models.User{
		Username:     username,
		PasswordHash: hashedPassword,
	}

	var token string
	errGroup := errgroup.Group{}
	errGroup.Go(func() error {
		return c.userStorage.SaveUser(ctx, user)
	})
	errGroup.Go(func() (err error) {
		token, err = c.jwtManager.Generate(user)
		return err
	})

	// no matter - error during saving user or generating
	if err := errGroup.Wait(); err != nil {
		c.logger.Error("got error during creation user: %v", err)
		return nil, status.Error(codes.Internal, "unexpected error during creation user")
	}

	return &servicepb.RegisterResponse{
		AccessToken: token,
	}, nil
}

func (c *Controller) Login(ctx context.Context, username string, password string) (*servicepb.LoginResponse, error) {
	var loginFailed = status.Error(codes.Unauthenticated, "User does not exists or password is incorrect")

	user, err := c.userStorage.GetUser(ctx, username)
	if err != nil {
		if errors.Is(err, models.ErrorNotExists) {
			return nil, loginFailed
		}
		return nil, status.Error(codes.Internal, "unexpected error storage processing")
	}
	if IsCorrectPassword(user, password) {
		token, err := c.jwtManager.Generate(user)
		if err != nil {
			c.logger.Error("Could not generate token: %v", err)
			return nil, status.Error(codes.Internal, "unexpected error")
		}
		return &servicepb.LoginResponse{
			AccessToken: token,
		}, nil
	} else {
		return nil, loginFailed
	}
}
