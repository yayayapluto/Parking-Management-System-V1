package database

import (
	"gorm.io/gorm"
)

type TransactionFunc func(tx *gorm.DB) error

func ExecWithTransaction(db *gorm.DB, fn TransactionFunc) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
