package models

import (
	"database/sql"
	"time"
)

type Zone struct {
	ID                 uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Name               string         `gorm:"size:100;not null" json:"name" bson:"name" validate:"required,max=100"`
	ZoneTypeID         uint           `gorm:"not null" json:"zone_type_id" bson:"zone_type_id" validate:"required"`
	Location           string         `gorm:"size:255;not null" json:"location" bson:"location" validate:"required,max=255"`
	MaximumCapacity    int            `gorm:"not null" json:"maximum_capacity" bson:"maximum_capacity" validate:"required,min=1"`
	OccupiedCount      int            `gorm:"default:0" json:"occupied_count" bson:"occupied_count" validate:"-"`
	AvailableSlots     int            `gorm:"default:0" json:"available_slots" bson:"available_slots" validate:"-"`
	DefaultFee         float64        `gorm:"type:decimal(10,2);not null" json:"default_fee" bson:"default_fee" validate:"required,min=0"`
	Is24Hours          bool           `gorm:"default:false" json:"is_24_hours" bson:"is_24_hours" validate:"-"`
	OperatingHourStart sql.NullString `gorm:"type:time" json:"operating_hour_start" bson:"operating_hour_start" validate:"-"`
	OperatingHourEnd   sql.NullString `gorm:"type:time" json:"operating_hour_end" bson:"operating_hour_end" validate:"-"`
	IsActive           bool           `gorm:"default:true" json:"is_active" bson:"is_active" validate:"-"`
	IsInMaintenance    bool           `gorm:"default:false" json:"is_in_maintenance" bson:"is_in_maintenance" validate:"-"`
	LastExitTime       *time.Time     `json:"last_exit_time" bson:"last_exit_time" validate:"-"`
	Description        string         `gorm:"type:text" json:"description" bson:"description" validate:"max=1000"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	ZoneType ZoneType `gorm:"foreignKey:ZoneTypeID" json:"zone_type,omitempty" bson:"zone_type,omitempty" validate:"-"`
}

func (Zone) TableName() string { return "zones" }
