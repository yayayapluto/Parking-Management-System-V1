package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type UserRepository interface {
	BaseRepository[models.User]
}

type userRepository struct {
	BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}
