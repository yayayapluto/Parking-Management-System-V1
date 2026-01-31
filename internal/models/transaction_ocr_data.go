package models

import "time"

type TransactionOCRData struct {
	ID                 uint       `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID      uint       `gorm:"not null;uniqueIndex" json:"transaction_id" bson:"transaction_id" validate:"required"`
	EntryPhotoPath     string     `gorm:"size:500" json:"entry_photo_path" bson:"entry_photo_path" validate:"omitempty,max=500"`
	EntryPlateCropPath string     `gorm:"size:500" json:"entry_plate_crop_path" bson:"entry_plate_crop_path" validate:"omitempty,max=500"`
	EntryPlateDetected string     `gorm:"size:20" json:"entry_plate_detected" bson:"entry_plate_detected" validate:"omitempty,max=20"`
	EntryConfidence    float64    `gorm:"type:decimal(5,2);default:0" json:"entry_confidence" bson:"entry_confidence" validate:"min=0,max=100"`
	EntryOCRStatus     string     `gorm:"type:varchar(20);default:pending" json:"entry_ocr_status" bson:"entry_ocr_status" validate:"required,oneof=pending success failed manual_override"`
	EntryProcessedAt   *time.Time `json:"entry_processed_at" bson:"entry_processed_at" validate:"-"`
	ExitPhotoPath      string     `gorm:"size:500" json:"exit_photo_path" bson:"exit_photo_path" validate:"omitempty,max=500"`
	ExitPlateCropPath  string     `gorm:"size:500" json:"exit_plate_crop_path" bson:"exit_plate_crop_path" validate:"omitempty,max=500"`
	ExitPlateDetected  string     `gorm:"size:20" json:"exit_plate_detected" bson:"exit_plate_detected" validate:"omitempty,max=20"`
	ExitConfidence     float64    `gorm:"type:decimal(5,2);default:0" json:"exit_confidence" bson:"exit_confidence" validate:"min=0,max=100"`
	ExitOCRStatus      string     `gorm:"type:varchar(20);default:pending" json:"exit_ocr_status" bson:"exit_ocr_status" validate:"required,oneof=pending success failed manual_override"`
	ExitProcessedAt    *time.Time `json:"exit_processed_at" bson:"exit_processed_at" validate:"-"`
	HasMismatch        bool       `gorm:"default:false" json:"has_mismatch" bson:"has_mismatch" validate:"-"`
	MismatchReason     string     `gorm:"size:255" json:"mismatch_reason" bson:"mismatch_reason" validate:"omitempty,max=255"`
	CreatedAt          time.Time  `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`
	UpdatedAt          time.Time  `gorm:"autoUpdateTime" json:"updated_at" bson:"updated_at" validate:"-"`

	// Relations
	Transaction ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
}

func (TransactionOCRData) TableName() string { return "transaction_ocr_data" }
