package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepo_GetMessage(t *testing.T) {
	r := NewRepository()
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
			msg, err := r.GetMessage(ctx, tc.name)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedMsg, msg)
		})
	}
}
