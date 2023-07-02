package cliclient

import (
	"fmt"
	"os"
	"sync"

	clipb "github.com/ksusonic/gophkeeper/proto/cli"

	"google.golang.org/protobuf/proto"
)

type Storage struct {
	path  string
	mutex sync.Mutex

	*clipb.Storage
}

func NewStorage(path string, ignoreErrors bool) (*Storage, error) {
	storage := &Storage{path: path, Storage: nil}
	file, _ := os.ReadFile(path)
	if file != nil {
		value := &clipb.Storage{}
		err := proto.Unmarshal(file, value)
		if err != nil {
			if ignoreErrors {
				return storage, nil
			}
			return nil, fmt.Errorf("could not unmarshall storage proto: %v", err)
		}
		storage.Storage = value
		return storage, nil
	}
	return storage, nil
}

func (c *Storage) GetValue() *clipb.Storage {
	if c.Storage == nil {
		return &clipb.Storage{}
	}
	return c.Storage
}

func (c *Storage) SetValue(storage *clipb.Storage) {
	c.mutex.Lock()
	c.Storage = storage
	c.mutex.Unlock()
}

func (c *Storage) Save() error {
	c.mutex.Lock()
	marshal, err := proto.Marshal(c.Storage)
	c.mutex.Unlock()
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
