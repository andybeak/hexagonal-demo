package ports

import (
	"github.com/andybeak/hexagonal-demo/internal/core/domain"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	GetUserById(id string) (domain.User, error)
}
