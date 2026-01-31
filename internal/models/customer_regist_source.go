package models

import "time"

type CustomerRegistSource struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Name        string    `gorm:"size:50;not null;uniqueIndex" json:"name" bson:"name" validate:"required,max=50"`
	Description string    `gorm:"type:text" json:"description" bson:"description" validate:"max=1000"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`
}

func (CustomerRegistSource) TableName() string { return "customer_regist_sources" }
