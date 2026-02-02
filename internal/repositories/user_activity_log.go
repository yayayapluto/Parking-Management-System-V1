package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type UserActivityLogRepository interface {
	BaseRepository[models.UserActivityLog]
}

type userActivityLogRepository struct {
	BaseRepository[models.UserActivityLog]
}

func NewUserActivityLogRepository(db *gorm.DB) UserActivityLogRepository {
	return &userActivityLogRepository{
		BaseRepository: NewBaseRepository[models.UserActivityLog](db),
	}
}
