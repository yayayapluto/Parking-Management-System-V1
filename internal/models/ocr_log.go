package models

import "time"

type OCRLog struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	TransactionID    uint      `gorm:"not null;index" json:"transaction_id" bson:"transaction_id" validate:"required"`
	OCRType          string    `gorm:"type:varchar(20);not null" json:"ocr_type" bson:"ocr_type" validate:"required,oneof=entry exit verification"`
	ImagePath        string    `gorm:"size:500;not null" json:"image_path" bson:"image_path" validate:"required,max=500"`
	RawOCRResult     string    `gorm:"type:text" json:"raw_ocr_result" bson:"raw_ocr_result" validate:"-"`
	DetectedText     string    `gorm:"size:255" json:"detected_text" bson:"detected_text" validate:"omitempty,max=255"`
	ValidatedPlate   string    `gorm:"size:20" json:"validated_plate" bson:"validated_plate" validate:"omitempty,max=20"`
	ConfidenceScore  float64   `gorm:"type:decimal(5,2);default:0" json:"confidence_score" bson:"confidence_score" validate:"min=0,max=100"`
	ProcessingTimeMs int       `gorm:"default:0" json:"processing_time_ms" bson:"processing_time_ms" validate:"min=0"`
	Status           string    `gorm:"type:varchar(20);not null;default:pending" json:"status" bson:"status" validate:"required,oneof=pending success failed timeout"`
	ErrorMessage     string    `gorm:"size:500" json:"error_message" bson:"error_message" validate:"omitempty,max=500"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at" bson:"created_at" validate:"-"`

	// Relations
	Transaction ParkingTransaction `gorm:"foreignKey:TransactionID" json:"transaction,omitempty" bson:"transaction,omitempty" validate:"-"`
}

func (OCRLog) TableName() string { return "ocr_logs" }
