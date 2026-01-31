package models

import "time"

type LostTicketFee struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID uint      `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	OriginalFee   float64   `gorm:"type:decimal(10,2);not null" json:"original_fee" bson:"original_fee" validate:"required,min=0"`
	LostTicketFee float64   `gorm:"type:decimal(10,2);not null" json:"lost_ticket_fee" bson:"lost_ticket_fee" validate:"required,min=0"`
	TotalCharged  float64   `gorm:"type:decimal(10,2);not null" json:"total_charged" bson:"total_charged" validate:"required,min=0"`
	ProcessedBy   uint      `gorm:"not null" json:"processed_by" bson:"processed_by" validate:"required"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Transaction     ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	ProcessedByUser User               `gorm:"foreignKey:ProcessedBy" json:"processed_by_user,omitempty" bson:"processed_by_user,omitempty" validate:"-"`
}

func (LostTicketFee) TableName() string { return "lost_ticket_fees" }
