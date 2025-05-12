package repo

import (
	"context"
	"errors"

	"go.uber.org/fx"
)

// Module provides the Repository implementation to an Fx application.
var Module = fx.Module(
	"repository",
	fx.Provide(NewRepository),
)

// ErrNotFound indicates a requested entity was not found in the repository.
var ErrNotFound = errors.New("repository: not found")

// Repository defines data access methods for domain entities.
type Repository interface {
	// GetMessage retrieves a message by name.
	// Returns ErrNotFound if the key does not exist.
	GetMessage(ctx context.Context, name string) (string, error)
}

// InMemoryRepo is a trivial in-memory repository implementation.
type InMemoryRepo struct {
	data map[string]string
}

// NewRepository constructs an InMemoryRepo with default data.
func NewRepository() Repository {
	return &InMemoryRepo{
		data: map[string]string{
			"world": "Hello, World!",
			"":      "Hello, Anonymous!",
		},
	}
}

// Ensure InMemoryRepo implements Repository.
var _ Repository = (*InMemoryRepo)(nil)

// GetMessage returns the message for the given name or ErrNotFound.
func (r *InMemoryRepo) GetMessage(ctx context.Context, name string) (string, error) {
	if name == "" {
		return r.data[name], nil
	}

	if msg, ok := r.data[name]; ok {
		return msg, nil
	}

	return "", ErrNotFound
}
