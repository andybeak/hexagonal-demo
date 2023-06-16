package ports

import (
	"context"
	"github.com/andybeak/hexagonal-demo/internal/core/domain"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, name string) (domain.User, error)
	GetUserById(ctx context.Context, id string) (domain.User, error)
}
