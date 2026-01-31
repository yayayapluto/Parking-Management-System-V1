package models

import (
	"github.com/lib/pq"
	"time"
)

type ReportCache struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id" bson:"id" validate:"-"`
	ReportType      string         `gorm:"type:varchar(20);not null;index" json:"report_type" bson:"report_type" validate:"required,oneof=daily weekly monthly zone_summary revenue occupancy"`
	ReportDate      time.Time      `gorm:"type:date;not null;index" json:"report_date" bson:"report_date" validate:"required"`
	ZoneID          *uint          `gorm:"index" json:"zone_id" bson:"zone_id" validate:"-"`
	ReportData      pq.StringArray `gorm:"type:jsonb" json:"report_data" bson:"report_data" validate:"-"`
	TotalEntries    int            `gorm:"default:0" json:"total_entries" bson:"total_entries" validate:"min=0"`
	TotalExits      int            `gorm:"default:0" json:"total_exits" bson:"total_exits" validate:"min=0"`
	TotalActive     int            `gorm:"default:0" json:"total_active" bson:"total_active" validate:"min=0"`
	TotalRevenue    float64        `gorm:"type:decimal(12,2);default:0" json:"total_revenue" bson:"total_revenue" validate:"min=0"`
	AverageDuration float64        `gorm:"type:decimal(8,2);default:0" json:"average_duration" bson:"average_duration" validate:"min=0"`
	PeakOccupancy   int            `gorm:"default:0" json:"peak_occupancy" bson:"peak_occupancy" validate:"min=0"`
	GeneratedAt     time.Time      `gorm:"not null" json:"generated_at" bson:"generated_at" validate:"required"`
	ExpiresAt       time.Time      `gorm:"not null;index" json:"expires_at" bson:"expires_at" validate:"required"`

	// Relations
	Zone *Zone `gorm:"foreignKey:ZoneID" json:"zone,omitempty" bson:"zone,omitempty" validate:"-"`
}

func (ReportCache) TableName() string { return "reports_cache" }
