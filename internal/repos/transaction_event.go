package repos

import (
	"gorm.io/gorm"
	"parking-management-system-v1/internal/models"
)

type TransactionEventRepository interface {
	BaseRepository[models.TransactionEvent]
}

type transactionEventRepository struct {
	BaseRepository[models.TransactionEvent]
}

func NewTransactionEventRepository(db *gorm.DB) TransactionEventRepository {
	return &transactionEventRepository{
		BaseRepository: NewBaseRepository[models.TransactionEvent](db),
	}
}
