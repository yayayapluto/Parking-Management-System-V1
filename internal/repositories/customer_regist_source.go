package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type CustomerRegistSourceRepository interface {
	BaseRepository[models.CustomerRegistSource]
}

type customerRegistSourceRepository struct {
	BaseRepository[models.CustomerRegistSource]
}

func NewCustomerRegistSourceRepository(db *gorm.DB) CustomerRegistSourceRepository {
	return &customerRegistSourceRepository{
		BaseRepository: NewBaseRepository[models.CustomerRegistSource](db),
	}
}
