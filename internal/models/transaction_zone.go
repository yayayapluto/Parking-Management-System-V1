package models

import "time"

type TransactionZone struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID   uint       `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	SuggestedZoneID uint       `gorm:"not null" json:"suggested_zone_id" bson:"suggested_zone_id" validate:"required"`
	ActualZoneID    uint       `gorm:"not null" json:"actual_zone_id" bson:"actual_zone_id" validate:"required"`
	EntryTime       time.Time  `gorm:"not null" json:"entry_time" bson:"entry_time" validate:"required"`
	ExitTime        *time.Time `json:"exit_time" bson:"exit_time" validate:"-"`
	DurationMinutes int        `gorm:"default:0" json:"duration_minutes" bson:"duration_minutes" validate:"-"`
	ZoneFee         float64    `gorm:"type:decimal(10,2);default:0" json:"zone_fee" bson:"zone_fee" validate:"min=0"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Transaction   ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	SuggestedZone Zone               `gorm:"foreignKey:SuggestedZoneID" json:"suggested_zone,omitempty" bson:"suggested_zone,omitempty" validate:"-"`
	ActualZone    Zone               `gorm:"foreignKey:ActualZoneID" json:"actual_zone,omitempty" bson:"actual_zone,omitempty" validate:"-"`
}

func (TransactionZone) TableName() string { return "transaction_zones" }
