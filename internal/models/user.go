package models

import "time"

type User struct {
	ID                  uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	Username            string     `gorm:"size:50;not null;uniqueIndex" json:"username" bson:"username" validate:"required,max=50,alphanum"`
	FullName            string     `gorm:"size:100;not null" json:"full_name" bson:"full_name" validate:"required,max=100"`
	Email               string     `gorm:"size:255;not null;uniqueIndex" json:"email" bson:"email" validate:"required,email,max=255"`
	Phone               string     `gorm:"size:20" json:"phone" bson:"phone" validate:"omitempty,max=20"`
	Password            string     `gorm:"size:255;not null" json:"-" bson:"password" validate:"required,min=8"`
	RoleID              uint       `gorm:"not null" json:"role_id" bson:"role_id" validate:"required"`
	IsActive            bool       `gorm:"default:true" json:"is_active" bson:"is_active" validate:"-"`
	IsLocked            bool       `gorm:"default:false" json:"is_locked" bson:"is_locked" validate:"-"`
	LastLoginAt         *time.Time `json:"last_login_at" bson:"last_login_at" validate:"-"`
	LastLoginIP         string     `gorm:"size:45" json:"last_login_ip" bson:"last_login_ip" validate:"-"`
	FailedLoginAttempts int        `gorm:"default:0" json:"failed_login_attempts" bson:"failed_login_attempts" validate:"-"`
	LockedUntil         *time.Time `json:"locked_until" bson:"locked_until" validate:"-"`
	CreatedAt           time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt           time.Time  `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Role Role `gorm:"foreignKey:RoleID" json:"role,omitempty" bson:"role,omitempty" validate:"-"`
}

func (User) TableName() string { return "users" }
