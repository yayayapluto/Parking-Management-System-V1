package models

import "time"

type Role struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Name         string    `gorm:"size:50;not null;uniqueIndex" json:"name" bson:"name" validate:"required,max=50"`
	PermissionID uint      `gorm:"not null" json:"permission_id" bson:"permission_id" validate:"required"`
	CreatedBy    string    `gorm:"size:100" json:"created_by" bson:"created_by" validate:"max=100"`
	UpdatedBy    string    `gorm:"size:100" json:"updated_by" bson:"updated_by" validate:"max=100"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Permission Permission `gorm:"foreignKey:PermissionID" json:"permission,omitempty" bson:"permission,omitempty" validate:"-"`
}

func (Role) TableName() string { return "roles" }
