package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ZoneTypeRepository interface {
	BaseRepository[models.ZoneType]
}

type zoneTypeRepository struct {
	BaseRepository[models.ZoneType]
}

func NewZoneTypeRepository(db *gorm.DB) ZoneTypeRepository {
	return &zoneTypeRepository{
		BaseRepository: NewBaseRepository[models.ZoneType](db),
	}
}
