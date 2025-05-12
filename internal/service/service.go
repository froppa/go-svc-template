package service

import (
	"context"
	"errors"

	"github.com/froppa/go-svc-template/internal/repo"
	"go.uber.org/fx"
)

// ErrNotFound indicates the service could not find requested data.
var ErrNotFound = errors.New("not found")

// Service defines business logic operations.
type Service interface {
	// Hello returns a greeting for the given name.
	Hello(ctx context.Context, name string) (string, error)
}

// serviceImpl is a concrete Service that uses a Repository.
type serviceImpl struct {
	repo repo.Repository
}

// Module provides the Service implementation to an Fx application.
var Module = fx.Module(
	"service",
	fx.Provide(NewService),
)

// NewService constructs a Service, injecting the Repository dependency.
func NewService(r repo.Repository) Service {
	return &serviceImpl{repo: r}
}

// Hello generates a greeting message using the repository.
// Propagates ErrNotFound if the repo xhas no entry for name.
func (s *serviceImpl) Hello(ctx context.Context, name string) (string, error) {
	return s.repo.GetMessage(ctx, name)
}
