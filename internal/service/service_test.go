package service

import (
	"context"
	"testing"

	"github.com/froppa/go-svc-template/internal/repo"
	"github.com/stretchr/testify/assert"
)

func TestService_Hello(t *testing.T) {
	r := repo.NewRepository()
	svc := NewService(r)
	ctx := context.Background()

	tests := []struct {
		name        string
		expectedMsg string
		expectErr   bool
	}{
		{"world", "Hello, World!", false},
		{"", "Hello, Anonymous!", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg, err := svc.Hello(ctx, tc.name)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedMsg, msg)
		})
	}
}
