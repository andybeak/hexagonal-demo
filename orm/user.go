package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// User can be a project entity, or just a DTO for persistence
type User struct {
	Model
	// Notice that we include persistence details in the struct that is external to core
	Name string `gorm:"type:varchar(32)"`
}
