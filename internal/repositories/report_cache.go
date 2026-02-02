package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ReportCacheRepository interface {
	BaseRepository[models.ReportCache]
}

type reportCacheRepository struct {
	BaseRepository[models.ReportCache]
}

func NewReportCacheRepository(db *gorm.DB) ReportCacheRepository {
	return &reportCacheRepository{
		BaseRepository: NewBaseRepository[models.ReportCache](db),
	}
}
