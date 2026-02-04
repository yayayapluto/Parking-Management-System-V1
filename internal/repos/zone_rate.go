package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ZoneRateRepository interface {
	BaseRepository[models.ZoneRate]
}

type zoneRateRepository struct {
	BaseRepository[models.ZoneRate]
}

func NewZoneRateRepository(db *gorm.DB) ZoneRateRepository {
	return &zoneRateRepository{
		BaseRepository: NewBaseRepository[models.ZoneRate](db),
	}
}
