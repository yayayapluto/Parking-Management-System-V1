package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ZoneRepository interface {
	BaseRepository[models.Zone]
}

type zoneRepository struct {
	BaseRepository[models.Zone]
}

func NewZoneRepository(db *gorm.DB) ZoneRepository {
	return &zoneRepository{
		BaseRepository: NewBaseRepository[models.Zone](db),
	}
}
