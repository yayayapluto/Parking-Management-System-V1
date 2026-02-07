package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type ParkingTransactionRepository interface {
	BaseRepository[models.ParkingTransaction]
	// GetDB returns the underlying GORM database instance for direct queries
	GetDB() *gorm.DB
}

type parkingTransactionRepository struct {
	BaseRepository[models.ParkingTransaction]
	db *gorm.DB
}

func NewParkingTransactionRepository(db *gorm.DB) ParkingTransactionRepository {
	return &parkingTransactionRepository{
		BaseRepository: NewBaseRepository[models.ParkingTransaction](db),
		db:             db,
	}
}

// GetDB returns the underlying GORM database instance for direct queries
func (r *parkingTransactionRepository) GetDB() *gorm.DB {
	return r.db
}
