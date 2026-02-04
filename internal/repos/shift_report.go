package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ShiftReportRepository interface {
	BaseRepository[models.ShiftReport]
}

type shiftReportRepository struct {
	BaseRepository[models.ShiftReport]
}

func NewShiftReportRepository(db *gorm.DB) ShiftReportRepository {
	return &shiftReportRepository{
		BaseRepository: NewBaseRepository[models.ShiftReport](db),
	}
}
