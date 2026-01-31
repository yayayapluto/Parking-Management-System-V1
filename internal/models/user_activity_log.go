package models

import (
	"github.com/lib/pq"
	"time"
)

type UserActivityLog struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	UserID      uint           `gorm:"not null;index" json:"user_id" bson:"user_id" validate:"required"`
	Action      string         `gorm:"size:50;not null" json:"action" bson:"action" validate:"required,max=50"`
	Module      string         `gorm:"size:50;not null" json:"module" bson:"module" validate:"required,max=50"`
	Description string         `gorm:"type:text" json:"description" bson:"description" validate:"max=5000"`
	IPAddress   string         `gorm:"size:45" json:"ip_address" bson:"ip_address" validate:"omitempty,max=45"`
	UserAgent   string         `gorm:"size:500" json:"user_agent" bson:"user_agent" validate:"omitempty,max=500"`
	RequestData pq.StringArray `gorm:"type:jsonb" json:"request_data" bson:"request_data" validate:"-"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user,omitempty" bson:"user,omitempty" validate:"-"`
}

func (UserActivityLog) TableName() string { return "user_activity_logs" }
