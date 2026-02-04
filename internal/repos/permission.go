package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type PermissionRepository interface {
	BaseRepository[models.Permission]
}

type permissionRepository struct {
	BaseRepository[models.Permission]
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{
		BaseRepository: NewBaseRepository[models.Permission](db),
	}
}
