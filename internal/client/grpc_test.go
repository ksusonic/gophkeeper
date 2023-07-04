package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGrpcClient(t *testing.T) {
	t.Run("non-directory", func(t *testing.T) {
		got, err := NewGrpcClient(":3200", "/dev/null")
		assert.Error(t, err)
		assert.ErrorContains(t, err, "not a directory")
		assert.Nil(t, got)
	})
	t.Run("non-exist", func(t *testing.T) {
		got, err := NewGrpcClient(":3200", "/tmp")
		assert.Error(t, err)
		assert.ErrorContains(t, err, "no such file or directory")
		assert.Nil(t, got)
	})
	t.Run("no-error", func(t *testing.T) {
		got, err := NewGrpcClient(":3200", "../../cert/server-tls")
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
