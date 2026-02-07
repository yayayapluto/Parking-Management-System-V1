package models

import "time"

type Permission struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Name        string    `gorm:"size:100;not null;uniqueIndex" json:"name" bson:"name" validate:"required,max=100"`
	Module      string    `gorm:"size:100" json:"module" bson:"module" validate:"omitempty,max=100"`
	Description string    `gorm:"type:text" json:"description" bson:"description" validate:"max=1000"`
	CreatedBy   string    `gorm:"size:100" json:"created_by" bson:"created_by" validate:"max=100"`
	UpdatedBy   string    `gorm:"size:100" json:"updated_by" bson:"updated_by" validate:"max=100"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`
}

func (Permission) TableName() string { return "permissions" }
