package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type PaymentMethodRepository interface {
	BaseRepository[models.PaymentMethod]
}

type paymentMethodRepository struct {
	BaseRepository[models.PaymentMethod]
}

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		BaseRepository: NewBaseRepository[models.PaymentMethod](db),
	}
}
