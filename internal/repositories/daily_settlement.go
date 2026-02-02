package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type DailySettlementRepository interface {
	BaseRepository[models.DailySettlement]
}

type dailySettlementRepository struct {
	BaseRepository[models.DailySettlement]
}

func NewDailySettlementRepository(db *gorm.DB) DailySettlementRepository {
	return &dailySettlementRepository{
		BaseRepository: NewBaseRepository[models.DailySettlement](db),
	}
}
