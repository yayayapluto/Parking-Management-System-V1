package models

import "time"

type DailySettlement struct {
	ID                uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Date              time.Time  `gorm:"type:date;not null;uniqueIndex" json:"date" bson:"date" validate:"required"`
	TotalTransactions int        `gorm:"default:0" json:"total_transactions" bson:"total_transactions" validate:"min=0"`
	TotalRevenue      float64    `gorm:"type:decimal(12,2);default:0" json:"total_revenue" bson:"total_revenue" validate:"min=0"`
	TotalCash         float64    `gorm:"type:decimal(12,2);default:0" json:"total_cash" bson:"total_cash" validate:"min=0"`
	TotalQRIS         float64    `gorm:"type:decimal(12,2);default:0" json:"total_qris" bson:"total_qris" validate:"min=0"`
	TotalRefunds      float64    `gorm:"type:decimal(12,2);default:0" json:"total_refunds" bson:"total_refunds" validate:"min=0"`
	TotalLostTickets  int        `gorm:"default:0" json:"total_lost_tickets" bson:"total_lost_tickets" validate:"min=0"`
	Variance          float64    `gorm:"type:decimal(12,2);default:0" json:"variance" bson:"variance" validate:"-"`
	ReconciledBy      *uint      `json:"reconciled_by" bson:"reconciled_by" validate:"-"`
	ReconciledAt      *time.Time `json:"reconciled_at" bson:"reconciled_at" validate:"-"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	ReconciledByUser *User `gorm:"foreignKey:ReconciledBy" json:"reconciled_by_user,omitempty" bson:"reconciled_by_user,omitempty" validate:"-"`
}

func (DailySettlement) TableName() string { return "daily_settlements" }
