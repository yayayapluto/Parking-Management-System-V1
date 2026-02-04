package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ParkingTransactionRepository interface {
	BaseRepository[models.ParkingTransaction]
}

type parkingTransactionRepository struct {
	BaseRepository[models.ParkingTransaction]
}

func NewParkingTransactionRepository(db *gorm.DB) ParkingTransactionRepository {
	return &parkingTransactionRepository{
		BaseRepository: NewBaseRepository[models.ParkingTransaction](db),
	}
}
