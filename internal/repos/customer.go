package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type CustomerRepository interface {
	BaseRepository[models.Customer]
}

type customerRepository struct {
	BaseRepository[models.Customer]
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		BaseRepository: NewBaseRepository[models.Customer](db),
	}
}
