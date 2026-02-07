package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"time"
)

// FeeCalculationResult holds the result of fee calculation
type FeeCalculationResult struct {
	DurationMinutes int
	DurationHours   float64
	HourlyRate      float64
	BaseFee         float64
	TotalFee        float64
}

// FeeCalculator defines the interface for calculating parking fees
type FeeCalculator interface {
	// CalculateFeeWithDB calculates the parking fee based on entry time, exit time, zone, and vehicle type
	CalculateFeeWithDB(ctx context.Context, entryTime time.Time, exitTime time.Time, zoneID uint, vehicleTypeID uint, db *gorm.DB) (FeeCalculationResult, error)
}

type feeCalculator struct {
	zoneRateRepo repos.ZoneRateRepository
}

// NewFeeCalculator creates a new instance of FeeCalculator
func NewFeeCalculator(zoneRateRepo repos.ZoneRateRepository) FeeCalculator {
	return &feeCalculator{
		zoneRateRepo: zoneRateRepo,
	}
}

// CalculateFeeWithDB calculates the parking fee based on entry time, exit time, zone, and vehicle type
// Logic:
// 1. Calculate duration in minutes
// 2. Fetch the hourly rate for the zone and vehicle type
// 3. Calculate the fee based on the hourly rate
// 4. Handle rounding: round up to the nearest hour (e.g., 2h 30m = 3h)
// 5. Apply daily max rate if applicable
func (fc *feeCalculator) CalculateFeeWithDB(
	ctx context.Context,
	entryTime time.Time,
	exitTime time.Time,
	zoneID uint,
	vehicleTypeID uint,
	db *gorm.DB,
) (FeeCalculationResult, error) {
	result := FeeCalculationResult{}

	// Validate input times
	if entryTime.After(exitTime) {
		return result, errors.New("entry time cannot be after exit time")
	}

	// Calculate duration in minutes
	duration := exitTime.Sub(entryTime)
	durationMinutes := int(duration.Minutes())
	result.DurationMinutes = durationMinutes

	// Fetch the zone rate for the given zone and vehicle type
	var zoneRate models.ZoneRate
	today := time.Now().Truncate(24 * time.Hour)

	// Query the zone rate from the database
	// This query finds the most applicable rate for the given zone and vehicle type
	// Criteria:
	// - zone_id = zoneID
	// - vehicle_type_id = vehicleTypeID
	// - is_active = true
	// - valid_from <= today AND valid_to >= today
	// - Order by effective_hour_from to get the most specific time range
	if err := db.WithContext(ctx).
		Where("zone_id = ? AND vehicle_type_id = ? AND is_active = ?", zoneID, vehicleTypeID, true).
		Where("valid_from <= ? AND valid_to >= ?", today, today).
		Order("effective_hour_from DESC").
		First(&zoneRate).Error; err != nil {
		// If no specific rate found, try to get the default rate for the zone
		if err := db.WithContext(ctx).
			Where("zone_id = ? AND vehicle_type_id = ? AND is_active = ?", zoneID, vehicleTypeID, true).
			Order("valid_from DESC").
			First(&zoneRate).Error; err != nil {
			return result, fmt.Errorf("no applicable rate found for zone %d and vehicle type %d", zoneID, vehicleTypeID)
		}
	}

	// Fee Calculation: Subtract FreeMinutes from the total duration before applying the hourly rate
	chargeableMinutes := durationMinutes - zoneRate.FreeMinutes
	if chargeableMinutes < 0 {
		chargeableMinutes = 0
	}

	// Calculate duration in hours (with rounding up)
	// Example: 2h 30m = 150 minutes = 2.5 hours -> rounds up to 3 hours
	chargeableDuration := time.Duration(chargeableMinutes) * time.Minute
	durationHours := math.Ceil(chargeableDuration.Hours())
	if durationHours < 1 && chargeableMinutes > 0 {
		durationHours = 1 // Minimum 1 hour charge if there's any chargeable time
	}
	result.DurationHours = durationHours

	result.HourlyRate = zoneRate.HourlyRate
	result.BaseFee = zoneRate.HourlyRate * durationHours

	// Check if there's a daily max rate and apply it if necessary
	if zoneRate.DailyMaxRate > 0 && result.BaseFee > zoneRate.DailyMaxRate {
		result.TotalFee = zoneRate.DailyMaxRate
	} else {
		result.TotalFee = result.BaseFee
	}

	return result, nil
}
