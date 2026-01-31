package models

import (
	"database/sql"
	"time"
)

type ZoneRate struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	ZoneID            uint           `gorm:"not null;index" json:"zone_id" bson:"zone_id" validate:"required"`
	VehicleTypeID     uint           `gorm:"not null;index" json:"vehicle_type_id" bson:"vehicle_type_id" validate:"required"`
	HourlyRate        float64        `gorm:"type:decimal(10,2);not null" json:"hourly_rate" bson:"hourly_rate" validate:"required,min=0"`
	DailyMaxRate      float64        `gorm:"type:decimal(10,2);not null" json:"daily_max_rate" bson:"daily_max_rate" validate:"required,min=0"`
	FreeMinutes       int            `gorm:"default:0" json:"free_minutes" bson:"free_minutes" validate:"min=0"`
	IsWeekend         bool           `gorm:"default:false" json:"is_weekend" bson:"is_weekend" validate:"-"`
	IsHoliday         bool           `gorm:"default:false" json:"is_holiday" bson:"is_holiday" validate:"-"`
	HolidayID         *uint          `gorm:"index" json:"holiday_id" bson:"holiday_id" validate:"-"`
	ValidFrom         time.Time      `gorm:"type:date;not null" json:"valid_from" bson:"valid_from" validate:"required"`
	ValidTo           time.Time      `gorm:"type:date;not null" json:"valid_to" bson:"valid_to" validate:"required"`
	EffectiveHourFrom sql.NullString `gorm:"type:time" json:"effective_hour_from" bson:"effective_hour_from" validate:"-"`
	EffectiveHourTo   sql.NullString `gorm:"type:time" json:"effective_hour_to" bson:"effective_hour_to" validate:"-"`
	IsActive          bool           `gorm:"default:true" json:"is_active" bson:"is_active" validate:"-"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Zone        Zone        `gorm:"foreignKey:ZoneID" json:"zone,omitempty" bson:"zone,omitempty" validate:"-"`
	VehicleType VehicleType `gorm:"foreignKey:VehicleTypeID" json:"vehicle_type,omitempty" bson:"vehicle_type,omitempty" validate:"-"`
	Holiday     *Holiday    `gorm:"foreignKey:HolidayID" json:"holiday,omitempty" bson:"holiday,omitempty" validate:"-"`
}

func (ZoneRate) TableName() string { return "zone_rates" }
