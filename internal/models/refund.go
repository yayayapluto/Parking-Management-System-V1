package models

import "time"

type Refund struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	PaymentID     uint       `gorm:"not null;index" json:"payment_id" bson:"payment_id" validate:"required"`
	TransactionID uint       `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	Amount        float64    `gorm:"type:decimal(10,2);not null" json:"amount" bson:"amount" validate:"required,min=0"`
	Reason        string     `gorm:"size:255;not null" json:"reason" bson:"reason" validate:"required,max=255"`
	RefundedBy    uint       `gorm:"not null" json:"refunded_by" bson:"refunded_by" validate:"required"`
	ApprovedBy    *uint      `json:"approved_by" bson:"approved_by" validate:"-"`
	Status        string     `gorm:"type:varchar(20);not null;default:requested" json:"status" bson:"status" validate:"required,oneof=requested pending_approval approved rejected completed"`
	RefundedAt    *time.Time `json:"refunded_at" bson:"refunded_at" validate:"-"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Payment        Payment            `gorm:"foreignKey:PaymentID" json:"payment,omitempty" bson:"payment,omitempty" validate:"-"`
	Transaction    ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	RefundedByUser User               `gorm:"foreignKey:RefundedBy" json:"refunded_by_user,omitempty" bson:"refunded_by_user,omitempty" validate:"-"`
	ApprovedByUser *User              `gorm:"foreignKey:ApprovedBy" json:"approved_by_user,omitempty" bson:"approved_by_user,omitempty" validate:"-"`
}

func (Refund) TableName() string { return "refunds" }
