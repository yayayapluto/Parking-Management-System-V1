package models

import (
	"github.com/lib/pq"
	"time"
)

type PaymentMethod struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Code      string         `gorm:"size:20;not null;uniqueIndex" json:"code" bson:"code" validate:"required,max=20,alphanum"`
	Name      string         `gorm:"size:50;not null" json:"name" bson:"name" validate:"required,max=50"`
	Config    pq.StringArray `gorm:"type:jsonb" json:"config" bson:"config" validate:"-"`
	IsActive  bool           `gorm:"default:true" json:"is_active" bson:"is_active" validate:"-"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`
}

func (PaymentMethod) TableName() string { return "payment_methods" }
