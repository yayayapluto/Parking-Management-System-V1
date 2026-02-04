package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type OCRLogRepository interface {
	BaseRepository[models.OCRLog]
}

type oCRLogRepository struct {
	BaseRepository[models.OCRLog]
}

func NewOCRLogRepository(db *gorm.DB) OCRLogRepository {
	return &oCRLogRepository{
		BaseRepository: NewBaseRepository[models.OCRLog](db),
	}
}
