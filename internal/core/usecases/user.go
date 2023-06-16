package usecases

import (
	"context"
	"github.com/andybeak/hexagonal-demo/internal/core/domain"
	"github.com/andybeak/hexagonal-demo/internal/core/ports"
)

func ProvideUserUseCase(
	userRepository ports.UserRepository,
) ports.UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

// userUseCase implements ports.UserUseCase
type userUseCase struct {
	userRepository ports.UserRepository
}

func (u userUseCase) CreateUser(ctx context.Context, name string) (domain.User, error) {
	user := domain.NewUser(name)
	return u.userRepository.Save(user)
}

func (u userUseCase) GetUserById(ctx context.Context, id string) (domain.User, error) {
	return u.userRepository.GetUserById(id)
}
