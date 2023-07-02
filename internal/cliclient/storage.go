package cliclient

import (
	"fmt"
	"os"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"

	"google.golang.org/protobuf/proto"
)

type Storage struct {
	Value *clipb.Storage
	path  string
}

func NewStorage(path string, ignoreErrors bool) (*Storage, error) {
	storage := &Storage{
		Value: &clipb.Storage{},
		path:  path,
	}
	file, _ := os.ReadFile(path)
	if file != nil {
		value := &clipb.Storage{}
		err := proto.Unmarshal(file, value)
		if err != nil {
			if ignoreErrors {
				return storage, nil
			}
			return nil, fmt.Errorf("could not unmarshall proto: %v", err)
		}
		storage.Value = value
		return storage, nil
	}
	return storage, nil
}

func (c *Storage) Save() error {
	marshal, err := proto.Marshal(c.Value)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(c.path, os.O_CREATE|os.O_WRONLY, 0666)
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
