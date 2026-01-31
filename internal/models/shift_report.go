package models

import "time"

type ShiftReport struct {
	ID                uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	UserID            uint      `gorm:"not null;index" json:"user_id" bson:"user_id" validate:"required"`
	ShiftStart        time.Time `gorm:"not null" json:"shift_start" bson:"shift_start" validate:"required"`
	ShiftEnd          time.Time `gorm:"not null" json:"shift_end" bson:"shift_end" validate:"required"`
	OpeningCash       float64   `gorm:"type:decimal(12,2);default:0" json:"opening_cash" bson:"opening_cash" validate:"min=0"`
	ClosingCash       float64   `gorm:"type:decimal(12,2);default:0" json:"closing_cash" bson:"closing_cash" validate:"min=0"`
	TotalTransactions int       `gorm:"default:0" json:"total_transactions" bson:"total_transactions" validate:"min=0"`
	TotalCash         float64   `gorm:"type:decimal(12,2);default:0" json:"total_cash" bson:"total_cash" validate:"min=0"`
	TotalQRIS         float64   `gorm:"type:decimal(12,2);default:0" json:"total_qris" bson:"total_qris" validate:"min=0"`
	TotalRevenue      float64   `gorm:"type:decimal(12,2);default:0" json:"total_revenue" bson:"total_revenue" validate:"min=0"`
	Discrepancy       float64   `gorm:"type:decimal(12,2);default:0" json:"discrepancy" bson:"discrepancy" validate:"-"`
	Notes             string    `gorm:"type:text" json:"notes" bson:"notes" validate:"max=5000"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty" bson:"user,omitempty" validate:"-"`
}

func (ShiftReport) TableName() string { return "shift_reports" }
