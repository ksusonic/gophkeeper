package auth

import (
	"context"

	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	servicepb "github.com/ksusonic/gophkeeper/proto/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var loginFailed = status.Error(codes.Unauthenticated, "User does not exists or password is incorrect")

// Controller for user auth
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

// Register by login and password
func (c *Controller) Register(ctx context.Context, username, password string) (*servicepb.RegisterResponse, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty username or password")
	}

	if user, err := c.userStorage.GetUser(ctx, username); err != nil {
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

	err = c.userStorage.SaveUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not save user: %v", err)
	}

	token, err := c.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate access token: %v", err)
	}

	return &servicepb.RegisterResponse{
		AccessToken: token,
	}, nil
}

// Login by login and password
func (c *Controller) Login(ctx context.Context, username string, password string) (*servicepb.LoginResponse, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty username or password")
	}

	user, err := c.userStorage.GetUser(ctx, username)
	if err != nil {
		return nil, status.Error(codes.Internal, "unexpected error storage processing")
	} else if user == nil {
		return nil, loginFailed
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
