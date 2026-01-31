package models

import (
	"github.com/lib/pq"
	"time"
)

type Payment struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID   uint           `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	PaymentMethodID uint           `gorm:"not null;index" json:"payment_method_id" bson:"payment_method_id" validate:"required"`
	SubtotalAmount  float64        `gorm:"type:decimal(10,2);not null" json:"subtotal_amount" bson:"subtotal_amount" validate:"required,min=0"`
	FeeCharged      float64        `gorm:"type:decimal(10,2);default:0" json:"fee_charged" bson:"fee_charged" validate:"min=0"`
	FinalAmount     float64        `gorm:"type:decimal(10,2);not null" json:"final_amount" bson:"final_amount" validate:"required,min=0"`
	PaymentGateway  string         `gorm:"size:50" json:"payment_gateway" bson:"payment_gateway" validate:"omitempty,max=50"`
	ExternalID      string         `gorm:"size:255" json:"external_id" bson:"external_id" validate:"omitempty,max=255"`
	QRCodeURL       string         `gorm:"size:500" json:"qr_code_url" bson:"qr_code_url" validate:"omitempty,max=500"`
	QRString        string         `gorm:"type:text" json:"qr_string" bson:"qr_string" validate:"-"`
	Status          string         `gorm:"type:varchar(20);not null;default:pending;index" json:"status" bson:"status" validate:"required,oneof=pending paid failed expired refunded cancelled"`
	PaymentDetails  string         `gorm:"type:text" json:"payment_details" bson:"payment_details" validate:"-"`
	CallbackData    pq.StringArray `gorm:"type:jsonb" json:"callback_data" bson:"callback_data" validate:"-"`
	ReceiptNumber   string         `gorm:"size:100;index" json:"receipt_number" bson:"receipt_number" validate:"omitempty,max=100"`
	PaidAt          *time.Time     `json:"paid_at" bson:"paid_at" validate:"-"`
	ExpiredAt       *time.Time     `json:"expired_at" bson:"expired_at" validate:"-"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Transaction   ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	PaymentMethod PaymentMethod      `gorm:"foreignKey:PaymentMethodID" json:"payment_method,omitempty" bson:"payment_method,omitempty" validate:"-"`
}

func (Payment) TableName() string { return "payments" }
