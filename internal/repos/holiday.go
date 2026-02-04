package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type HolidayRepository interface {
	BaseRepository[models.Holiday]
}

type holidayRepository struct {
	BaseRepository[models.Holiday]
}

func NewHolidayRepository(db *gorm.DB) HolidayRepository {
	return &holidayRepository{
		BaseRepository: NewBaseRepository[models.Holiday](db),
	}
}
