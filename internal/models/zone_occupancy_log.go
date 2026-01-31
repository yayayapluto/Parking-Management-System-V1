package models

import "time"

type ZoneOccupancyLog struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	ZoneID         uint      `gorm:"not null;index" json:"zone_id" bson:"zone_id" validate:"required"`
	TransactionID  uint      `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	OperatorID     uint      `gorm:"not null;index" json:"operator_id" bson:"operator_id" validate:"required"`
	OccupiedCount  int       `gorm:"not null" json:"occupied_count" bson:"occupied_count" validate:"required,min=0"`
	AvailableSlots int       `gorm:"not null" json:"available_slots" bson:"available_slots" validate:"required,min=0"`
	EventType      string    `gorm:"type:varchar(20);not null" json:"event_type" bson:"event_type" validate:"required,oneof=entry exit manual_adjustment system_sync"`
	Notes          string    `gorm:"type:text" json:"notes" bson:"notes" validate:"max=5000"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Zone        Zone               `gorm:"foreignKey:ZoneID" json:"zone,omitempty" bson:"zone,omitempty" validate:"-"`
	Transaction ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	Operator    User               `gorm:"foreignKey:OperatorID" json:"operator,omitempty" bson:"operator,omitempty" validate:"-"`
}

func (ZoneOccupancyLog) TableName() string { return "zone_occupancy_logs" }
