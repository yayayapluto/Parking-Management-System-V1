package models

import "time"

type TransactionEvent struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID uint      `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	EventType     string    `gorm:"type:varchar(20);not null" json:"event_type" bson:"event_type" validate:"required,oneof=entry exit manual_entry manual_exit correction lost_ticket"`
	OperatorID    uint      `gorm:"not null;index" json:"operator_id" bson:"operator_id" validate:"required"`
	PhotoPath     string    `gorm:"size:500" json:"photo_path" bson:"photo_path" validate:"omitempty,max=500"`
	PlateDetected string    `gorm:"size:20" json:"plate_detected" bson:"plate_detected" validate:"omitempty,max=20"`
	RfidDetected  string    `gorm:"size:50" json:"rfid_detected" bson:"rfid_detected" validate:"omitempty,max=50"`
	Notes         string    `gorm:"type:text" json:"notes" bson:"notes" validate:"max=5000"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Transaction ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	Operator    User               `gorm:"foreignKey:OperatorID" json:"operator,omitempty" bson:"operator,omitempty" validate:"-"`
}

func (TransactionEvent) TableName() string { return "transaction_events" }
