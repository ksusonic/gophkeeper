package clihelper

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func TestNewStorage(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "simple create %1",
			path:    t.TempDir(),
			wantErr: false,
		},
		{
			name:    "simple create %2",
			path:    t.TempDir(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStorage(tt.path, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestStorage_LoginInterceptor(t *testing.T) {
	testToken := "super-token"
	expectedToken := "bearer " + testToken
	tests := []struct {
		name    string
		storage *clipb.Storage
		ctx     *cli.Context
		wantErr bool
	}{
		{
			name:    "ok",
			storage: &clipb.Storage{Token: testToken},
			ctx:     &cli.Context{Context: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{Storage: tt.storage}
			err := s.LoginInterceptor(tt.ctx)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				md, ok := metadata.FromOutgoingContext(tt.ctx.Context)
				assert.True(t, ok)
				assert.NotEmpty(t, md)
				assert.Len(t, md.Get("authorization"), 1)
				assert.Equal(t, md.Get("authorization")[0], expectedToken)
			}
		})
	}
}

func TestStorage_Save(t *testing.T) {
	tests := []struct {
		name    string
		storage *clipb.Storage
		path    string
		ctx     *cli.Context
		wantErr bool
	}{
		{
			name:    "saving ok",
			storage: &clipb.Storage{Token: "testToken"},
			path: func() string {
				tempFile, err := os.CreateTemp(t.TempDir(), "test-file")
				if err != nil {
					t.Fatal(err)
				}
				return tempFile.Name()
			}(),
			ctx:     &cli.Context{Context: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				path:    tt.path,
				Storage: tt.storage,
			}
			if err := s.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			file, err := os.Open(tt.path)
			if err != nil {
				t.Fatal(err)
			}
			all, err := io.ReadAll(file)
			assert.NoError(t, err)
			assert.NotEmpty(t, all)

			actual := &clipb.Storage{}
			assert.NoError(t, proto.Unmarshal(all, actual))
			assert.Equal(t, tt.storage.String(), actual.String())
		})
	}
}

func TestStorage_TokenIsValid(t *testing.T) {
	tests := []struct {
		name    string
		storage *clipb.Storage
		want    bool
	}{
		{
			"simple test",
			&clipb.Storage{Token: "woohoo"},
			true,
		},
		{
			"empty test",
			&clipb.Storage{Token: ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Storage: tt.storage,
			}
			if got := s.TokenIsValid(context.TODO()); got != tt.want {
				t.Errorf("TokenIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_TokenIsValidError(t *testing.T) {
	tests := []struct {
		name    string
		storage *clipb.Storage
		wantErr error
	}{
		{
			"simple test",
			&clipb.Storage{Token: "woohoo"},
			nil,
		},
		{
			"empty test",
			&clipb.Storage{Token: ""},
			fmt.Errorf("access token is empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Storage: tt.storage,
			}
			err := s.TokenIsValidError(context.TODO())
			if tt.wantErr != nil {
				assert.ErrorContains(t, err, tt.wantErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
