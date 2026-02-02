package repositories

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type TransactionZoneRepository interface {
	BaseRepository[models.TransactionZone]
}

type transactionZoneRepository struct {
	BaseRepository[models.TransactionZone]
}

func NewTransactionZoneRepository(db *gorm.DB) TransactionZoneRepository {
	return &transactionZoneRepository{
		BaseRepository: NewBaseRepository[models.TransactionZone](db),
	}
}
