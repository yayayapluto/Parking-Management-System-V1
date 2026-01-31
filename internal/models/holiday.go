package models

import "time"

type Holiday struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Date         time.Time `gorm:"type:date;not null;uniqueIndex" json:"date" bson:"date" validate:"required"`
	Name         string    `gorm:"size:100;not null" json:"name" bson:"name" validate:"required,max=100"`
	AffectsRates float64   `gorm:"type:decimal(5,2);default:1.0" json:"affects_rates" bson:"affects_rates" validate:"min=0"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`
}

func (Holiday) TableName() string { return "holidays" }
