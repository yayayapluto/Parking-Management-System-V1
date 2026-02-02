package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ManualCorrectionRepository interface {
	BaseRepository[models.ManualCorrection]
}

type manualCorrectionRepository struct {
	BaseRepository[models.ManualCorrection]
}

func NewManualCorrectionRepository(db *gorm.DB) ManualCorrectionRepository {
	return &manualCorrectionRepository{
		BaseRepository: NewBaseRepository[models.ManualCorrection](db),
	}
}
