package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type TransactionOCRDataRepository interface {
	BaseRepository[models.TransactionOCRData]
}

type transactionOCRDataRepository struct {
	BaseRepository[models.TransactionOCRData]
}

func NewTransactionOCRDataRepository(db *gorm.DB) TransactionOCRDataRepository {
	return &transactionOCRDataRepository{
		BaseRepository: NewBaseRepository[models.TransactionOCRData](db),
	}
}
