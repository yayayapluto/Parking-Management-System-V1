package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type LostTicketFeeRepository interface {
	BaseRepository[models.LostTicketFee]
}

type lostTicketFeeRepository struct {
	BaseRepository[models.LostTicketFee]
}

func NewLostTicketFeeRepository(db *gorm.DB) LostTicketFeeRepository {
	return &lostTicketFeeRepository{
		BaseRepository: NewBaseRepository[models.LostTicketFee](db),
	}
}
