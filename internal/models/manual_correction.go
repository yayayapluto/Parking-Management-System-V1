package models

import "time"

type ManualCorrection struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID  uint      `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	FieldCorrected string    `gorm:"size:100;not null" json:"field_corrected" bson:"field_corrected" validate:"required,max=100"`
	OldValue       string    `gorm:"size:255;not null" json:"old_value" bson:"old_value" validate:"required,max=255"`
	NewValue       string    `gorm:"size:255;not null" json:"new_value" bson:"new_value" validate:"required,max=255"`
	Reason         string    `gorm:"size:500;not null" json:"reason" bson:"reason" validate:"required,max=500"`
	CorrectedBy    uint      `gorm:"not null" json:"corrected_by" bson:"corrected_by" validate:"required"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Transaction     ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
	CorrectedByUser User               `gorm:"foreignKey:CorrectedBy" json:"corrected_by_user,omitempty" bson:"corrected_by_user,omitempty" validate:"-"`
}

func (ManualCorrection) TableName() string { return "manual_corrections" }
