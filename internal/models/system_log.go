package models

import "time"

type SystemLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Level      int       `gorm:"not null;index" json:"level" bson:"level" validate:"required,min=0,max=5"`
	Service    string    `gorm:"size:100;not null;index" json:"service" bson:"service" validate:"required,max=100"`
	Message    string    `gorm:"type:text;not null" json:"message" bson:"message" validate:"required"`
	Context    string    `gorm:"type:text" json:"context" bson:"context" validate:"-"`
	StackTrace string    `gorm:"type:text" json:"stack_trace" bson:"stack_trace" validate:"-"`
	CreatedAt  time.Time `gorm:"autoCreateTime;index" json:"created_at" bson:"created_at" validate:"-"`
}

func (SystemLog) TableName() string { return "system_logs" }
