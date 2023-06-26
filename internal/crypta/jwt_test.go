package crypta

import (
	"strings"
	"testing"
	"time"

	"github.com/ksusonic/gophkeeper/internal/config"
	"github.com/ksusonic/gophkeeper/internal/models"
	"github.com/stretchr/testify/assert"
)

const testSecretKey = "i-am-super-secret"

func TestJWTManager_Generate(t *testing.T) {

	type args struct {
		user *models.User
	}
	tests := []struct {
		name                string
		config              config.AuthConfig
		args                args
		expectedTokenPrefix string
	}{
		{
			name: "simple generation",
			config: config.AuthConfig{
				SaltKey:  testSecretKey,
				TokenTTL: time.Minute,
			},
			args: args{
				user: &models.User{
					Username:     "dandex",
					PasswordHash: "aaaa",
				},
			},
			expectedTokenPrefix: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manager := NewJWTManager(tt.config)
			got, err := manager.Generate(tt.args.user)
			assert.NoError(t, err, "Gor error from generation")
			assert.Equal(t, tt.expectedTokenPrefix, strings.SplitN(got, ".", 2)[0])
		})
	}
}
