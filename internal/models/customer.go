package models

import "time"

type Customer struct {
	ID                   uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	RfidUID              string     `gorm:"size:50;index" json:"rfid_uid" bson:"rfid_uid" validate:"omitempty,max=50"`
	Name                 string     `gorm:"size:100" json:"name" bson:"name" validate:"omitempty,max=100"`
	Phone                string     `gorm:"size:20" json:"phone" bson:"phone" validate:"omitempty,max=20"`
	RegistrationSourceID *uint      `gorm:"index" json:"registration_source_id" bson:"registration_source_id" validate:"-"`
	IsRegistered         bool       `gorm:"default:false" json:"is_registered" bson:"is_registered" validate:"-"`
	RegisteredAt         *time.Time `json:"registered_at" bson:"registered_at" validate:"-"`
	TotalVisits          int        `gorm:"default:0" json:"total_visits" bson:"total_visits" validate:"-"`
	TotalSpent           float64    `gorm:"type:decimal(12,2);default:0" json:"total_spent" bson:"total_spent" validate:"-"`
	CreatedAt            time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt            time.Time  `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	RegistrationSource *CustomerRegistSource `gorm:"foreignKey:RegistrationSourceID" json:"registration_source,omitempty" bson:"registration_source,omitempty" validate:"-"`
	Vehicles           []Vehicle             `gorm:"foreignKey:CustomerID" json:"vehicles,omitempty" bson:"vehicles,omitempty" validate:"-"`
}

func (Customer) TableName() string { return "customers" }
