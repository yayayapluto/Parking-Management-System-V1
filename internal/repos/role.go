package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type RoleRepository interface {
	BaseRepository[models.Role]
}

type roleRepository struct {
	BaseRepository[models.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		BaseRepository: NewBaseRepository[models.Role](db),
	}
}
