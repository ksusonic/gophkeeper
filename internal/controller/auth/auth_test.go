package auth

import (
	"context"
	"testing"
	"time"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/crypta"
	"github.com/ksusonic/gophkeeper/internal/logging"
	"github.com/ksusonic/gophkeeper/internal/models"
	"github.com/ksusonic/gophkeeper/test/teststorage"
	"github.com/mborders/logmatic"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type fields struct {
	userStorage UserStorage
	jwtManager  *crypta.JWTManager
	logger      logging.Logger
}

func newFields() fields {
	return fields{
		userStorage: teststorage.NewTestStorage(),
		jwtManager: crypta.NewJWTManager(config.AuthConfig{
			SaltKey:  "test!",
			TokenTTL: time.Minute * 5,
		}),
		logger: logmatic.NewLogger(),
	}
}

func newFieldsWithPreset(t *testing.T, users ...models.User) fields {
	f := newFields()
	for _, u := range users {
		err := f.userStorage.SaveUser(context.TODO(), &u)
		if err != nil {
			t.Fatalf("cannot init storage: %v", err)
		}
	}
	return f
}

func genPassword(t *testing.T, password string) string {
	hashPassword, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}
	return hashPassword
}

func TestController_Register(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "simple test #1",
			fields: newFields(),
			args: args{
				username: "ksusonic",
				password: "super-password",
			},
			wantErr: nil,
		},
		{
			name:   "empty password test #2",
			fields: newFields(),
			args: args{
				username: "ksusonic",
				password: "",
			},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
		{
			name:   "empty username test #2",
			fields: newFields(),
			args: args{
				username: "",
				password: "keks",
			},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
		{
			name:    "empty params test #2",
			fields:  newFields(),
			args:    args{},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
		{
			name: "already exists test #2",
			fields: newFieldsWithPreset(t, models.User{
				Username:     "ksusonic",
				PasswordHash: "123",
			}),
			args: args{
				username: "ksusonic",
				password: "321",
			},
			wantErr: status.Error(codes.AlreadyExists, "username ksusonic already taken"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				userStorage: tt.fields.userStorage,
				jwtManager:  tt.fields.jwtManager,
				logger:      tt.fields.logger,
			}
			got, err := c.Register(context.TODO(), tt.args.username, tt.args.password)
			if tt.wantErr != nil {
				assert.Error(t, err, "expected handler to give error")
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got.GetAccessToken(), "AccessToken is empty!")
			}
		})
	}
}

func TestController_Login(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "simple test #1",
			fields: newFieldsWithPreset(t, models.User{
				Username:     "ksusonic",
				PasswordHash: genPassword(t, "super-password"),
			}),
			args: args{
				username: "ksusonic",
				password: "super-password",
			},
			wantErr: nil,
		},
		{
			name: "wrong password test #2",
			fields: newFieldsWithPreset(t, models.User{
				Username:     "ksusonic",
				PasswordHash: "wrong",
			}),
			args: args{
				username: "ksusonic",
				password: "super-password",
			},
			wantErr: status.Error(codes.Unauthenticated, "User does not exists or password is incorrect"),
		},
		{
			name: "non-existing test #3",
			fields: newFieldsWithPreset(t, models.User{
				Username:     "admin",
				PasswordHash: "wrong",
			}),
			args: args{
				username: "ksusonic",
				password: "super-password",
			},
			wantErr: status.Error(codes.Unauthenticated, "User does not exists or password is incorrect"),
		},
		{
			name:   "empty password test #4",
			fields: newFields(),
			args: args{
				username: "ksusonic",
				password: "",
			},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
		{
			name:   "empty username test #5",
			fields: newFields(),
			args: args{
				username: "",
				password: "keks",
			},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
		{
			name:    "empty params test #6",
			fields:  newFields(),
			args:    args{},
			wantErr: status.Error(codes.InvalidArgument, "empty username or password"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				userStorage: tt.fields.userStorage,
				jwtManager:  tt.fields.jwtManager,
				logger:      tt.fields.logger,
			}
			got, err := c.Login(context.TODO(), tt.args.username, tt.args.password)
			if tt.wantErr != nil {
				assert.Error(t, err, "expected handler to give error")
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, got.GetAccessToken(), "AccessToken is empty!")
			}
		})
	}
}
