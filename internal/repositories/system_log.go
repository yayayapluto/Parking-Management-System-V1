package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type SystemLogRepository interface {
	BaseRepository[models.SystemLog]
}

type systemLogRepository struct {
	BaseRepository[models.SystemLog]
}

func NewSystemLogRepository(db *gorm.DB) SystemLogRepository {
	return &systemLogRepository{
		BaseRepository: NewBaseRepository[models.SystemLog](db),
	}
}
