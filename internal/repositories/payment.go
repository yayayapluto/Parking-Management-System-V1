package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type PaymentRepository interface {
	BaseRepository[models.Payment]
}

type paymentRepository struct {
	BaseRepository[models.Payment]
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		BaseRepository: NewBaseRepository[models.Payment](db),
	}
}
