package models

import "time"

type VehicleType struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Code        string    `gorm:"size:10;not null;uniqueIndex" json:"code" bson:"code" validate:"required,max=10,alphanum"`
	Name        string    `gorm:"size:50;not null" json:"name" bson:"name" validate:"required,max=50"`
	Description string    `gorm:"size:255" json:"description" bson:"description" validate:"max=255"`
	IsActive    bool      `gorm:"default:true" json:"is_active" bson:"is_active" validate:"-"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`
}

func (VehicleType) TableName() string { return "vehicle_types" }
