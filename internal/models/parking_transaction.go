package models

import "time"

type ParkingTransaction struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	CustomerID      *uint      `gorm:"index" json:"customer_id" bson:"customer_id" validate:"-"`
	OperatorID      uint       `gorm:"not null;index" json:"operator_id" bson:"operator_id" validate:"required"`
	VehicleTypeID   uint       `gorm:"not null;index" json:"vehicle_type_id" bson:"vehicle_type_id" validate:"required"`
	PlateNumber     string     `gorm:"size:20;not null;index" json:"plate_number" bson:"plate_number" validate:"required,max=20"`
	RfidUID         string     `gorm:"size:50;index" json:"rfid_uid" bson:"rfid_uid" validate:"omitempty,max=50"`
	EntryTime       time.Time  `gorm:"not null;index" json:"entry_time" bson:"entry_time" validate:"required"`
	ExitTime        *time.Time `gorm:"index" json:"exit_time" bson:"exit_time" validate:"-"`
	DurationMinutes int        `gorm:"default:0" json:"duration_minutes" bson:"duration_minutes" validate:"-"`
	BaseFee         float64    `gorm:"type:decimal(10,2);default:0" json:"base_fee" bson:"base_fee" validate:"min=0"`
	DiscountAmount  float64    `gorm:"type:decimal(10,2);default:0" json:"discount_amount" bson:"discount_amount" validate:"min=0"`
	AdditionalFees  float64    `gorm:"type:decimal(10,2);default:0" json:"additional_fees" bson:"additional_fees" validate:"min=0"`
	TotalFee        float64    `gorm:"type:decimal(10,2);default:0" json:"total_fee" bson:"total_fee" validate:"min=0"`
	Status          string     `gorm:"type:varchar(20);not null;default:active;index" json:"status" bson:"status" validate:"required,oneof=active completed cancelled expired lost_ticket"`
	PaymentStatus   string     `gorm:"type:varchar(20);not null;default:unpaid;index" json:"payment_status" bson:"payment_status" validate:"required,oneof=unpaid pending paid refunded failed"`
	Notes           string     `gorm:"type:text" json:"notes" bson:"notes" validate:"max=5000"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Customer         *Customer           `gorm:"foreignKey:CustomerID" json:"customer,omitempty" bson:"customer,omitempty" validate:"-"`
	Operator         User                `gorm:"foreignKey:OperatorID" json:"operator,omitempty" bson:"operator,omitempty" validate:"-"`
	VehicleType      VehicleType         `gorm:"foreignKey:VehicleTypeID" json:"vehicle_type,omitempty" bson:"vehicle_type,omitempty" validate:"-"`
	TransactionZones []TransactionZone   `gorm:"foreignKey:TransactionID" json:"transaction_zones,omitempty" bson:"transaction_zones,omitempty" validate:"-"`
	OCRData          *TransactionOCRData `gorm:"foreignKey:TransactionID" json:"ocr_data,omitempty" bson:"ocr_data,omitempty" validate:"-"`
	Events           []TransactionEvent  `gorm:"foreignKey:TransactionID" json:"events,omitempty" bson:"events,omitempty" validate:"-"`
	Payment          *Payment            `gorm:"foreignKey:TransactionID" json:"payment,omitempty" bson:"payment,omitempty" validate:"-"`
}

func (ParkingTransaction) TableName() string { return "parking_transactions" }
