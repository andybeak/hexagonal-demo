package adapters

import (
	"github.com/andybeak/hexagonal-demo/internal/core/domain"
	"github.com/andybeak/hexagonal-demo/internal/core/ports"
	"github.com/andybeak/hexagonal-demo/orm"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ProvideUserRepository(db *gorm.DB) ports.UserRepository {
	return &mySQLUserRepository{
		db: db,
	}
}

// mySQLUserRepository implements ports.UserRepository
type mySQLUserRepository struct {
	db *gorm.DB
}

func (u mySQLUserRepository) Save(user domain.User) (domain.User, error) {
	ormUser := orm.User{
		Model: orm.Model{
			ID: uuid.Must(uuid.Parse(user.Id)),
		},
		Name: user.Name,
	}
	if err := u.db.Create(&ormUser).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u mySQLUserRepository) GetUserById(id string) (domain.User, error) {
	var ormUser orm.User
	if err := u.db.First(&ormUser, "id = ?", id).Error; err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:   ormUser.ID.String(),
		Name: ormUser.Name,
	}, nil
}
