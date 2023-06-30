package teststorage

import (
	"context"
	"fmt"

	"github.com/ksusonic/gophkeeper/internal/models"
)

type TestStorage struct {
	usernameToUser map[string]models.User
	secrets        map[string]models.Secret
}

func NewTestStorage() *TestStorage {
	return &TestStorage{
		make(map[string]models.User),
		make(map[string]models.Secret),
	}
}

func (t *TestStorage) SaveUser(_ context.Context, user *models.User) error {
	if _, ok := t.usernameToUser[user.Username]; ok {
		return fmt.Errorf("user already exists")
	}
	t.usernameToUser[user.Username] = *user
	return nil
}

func (t *TestStorage) GetUser(_ context.Context, username string) (*models.User, error) {
	if user, ok := t.usernameToUser[username]; ok {
		u := user
		return &u, nil
	}
	return nil, nil
}
