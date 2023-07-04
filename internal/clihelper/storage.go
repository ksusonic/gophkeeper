package clihelper

import (
	"context"
	"fmt"
	"os"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"
	grpcMetadata "google.golang.org/grpc/metadata"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/proto"
)

type Storage struct {
	path string

	*clipb.Storage
}

func NewStorage(path string, ignoreErrors bool) (*Storage, error) {
	storage := &Storage{path: path, Storage: &clipb.Storage{}}
	file, _ := os.ReadFile(path)
	if file != nil {
		err := proto.Unmarshal(file, storage.Storage)
		if err != nil {
			if ignoreErrors {
				return storage, nil
			}
			return nil, fmt.Errorf("could not unmarshall storage proto: %v", err)
		}
		return storage, nil
	}
	return storage, nil
}

func (s *Storage) Save() error {
	marshal, err := proto.Marshal(s.Storage)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("could not create temp dir for file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) TokenIsValidError(_ context.Context) error {
	if len(s.Token) == 0 {
		return fmt.Errorf("access token is empty")
	}
	// TODO: validate handler
	return nil
}

func (s *Storage) TokenIsValid(ctx context.Context) bool {
	return s.TokenIsValidError(ctx) == nil
}

func (s *Storage) LoginInterceptor(ctx *cli.Context) error {
	if err := s.TokenIsValidError(ctx.Context); err != nil {
		return fmt.Errorf("you need to be logged in: %w", err)
	}

	md := grpcMetadata.Pairs("authorization", fmt.Sprintf("%s %v", "bearer", s.Token))
	ctx.Context = metadata.MD(md).ToOutgoing(ctx.Context)
	return nil
}
