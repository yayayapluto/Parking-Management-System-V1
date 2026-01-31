package models

import "time"

type Vehicle struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	CustomerID    *uint      `gorm:"index" json:"customer_id" bson:"customer_id" validate:"-"`
	VehicleTypeID uint       `gorm:"not null;index" json:"vehicle_type_id" bson:"vehicle_type_id" validate:"required"`
	PlateNumber   string     `gorm:"size:20;not null;uniqueIndex" json:"plate_number" bson:"plate_number" validate:"required,max=20"`
	Description   string     `gorm:"type:text" json:"description" bson:"description" validate:"max=1000"`
	LastSeenAt    *time.Time `json:"last_seen_at" bson:"last_seen_at" validate:"-"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Customer    *Customer   `gorm:"foreignKey:CustomerID" json:"customer,omitempty" bson:"customer,omitempty" validate:"-"`
	VehicleType VehicleType `gorm:"foreignKey:VehicleTypeID" json:"vehicle_type,omitempty" bson:"vehicle_type,omitempty" validate:"-"`
}

func (Vehicle) TableName() string { return "vehicles" }
