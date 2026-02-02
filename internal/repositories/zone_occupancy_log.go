package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ZoneOccupancyLogRepository interface {
	BaseRepository[models.ZoneOccupancyLog]
}

type zoneOccupancyLogRepository struct {
	BaseRepository[models.ZoneOccupancyLog]
}

func NewZoneOccupancyLogRepository(db *gorm.DB) ZoneOccupancyLogRepository {
	return &zoneOccupancyLogRepository{
		BaseRepository: NewBaseRepository[models.ZoneOccupancyLog](db),
	}
}
