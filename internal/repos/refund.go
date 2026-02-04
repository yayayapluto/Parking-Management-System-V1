package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type RefundRepository interface {
	BaseRepository[models.Refund]
}

type refundRepository struct {
	BaseRepository[models.Refund]
}

func NewRefundRepository(db *gorm.DB) RefundRepository {
	return &refundRepository{
		BaseRepository: NewBaseRepository[models.Refund](db),
	}
}
