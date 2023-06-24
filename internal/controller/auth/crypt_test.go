package auth

import (
	"testing"

	"github.com/ksusonic/gophkeeper/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestIsCorrectPassword(t *testing.T) {
	type args struct {
		user     *models.User
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple test on true",
			args: args{
				user: &models.User{
					Username:     "dandex",
					PasswordHash: "$2a$10$Q1iUtXzFwnBAME33MW7fLOMyBZYTGbvTivLFg951rGKgJcUHrRDAi",
				},
				password: "12345",
			},
			want: true,
		},
		{
			name: "simple test on false",
			args: args{
				user: &models.User{
					Username:     "dandex",
					PasswordHash: "",
				},
				password: "12345",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, IsCorrectPassword(tt.args.user, tt.args.password), "IsCorrectPassword(%v, %v)", tt.args.user, tt.args.password)
		})
	}
}
