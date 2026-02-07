package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"parking-management-system-v1/internal/dto/requests"
	"parking-management-system-v1/internal/dto/responses"
	"parking-management-system-v1/internal/models"
	"parking-management-system-v1/internal/repos"
	"parking-management-system-v1/pkg/helpers"
	"regexp"
	"strings"
	"time"
)

// ParkingTransactionService defines the business logic for parking transactions
type ParkingTransactionService interface {
	// RegisterEntry handles the vehicle entry (gate in) logic
	RegisterEntry(ctx context.Context, req requests.ParkingEntryRequest, operatorID uint) (responses.ParkingEntryResponse, error)
}

type parkingTransactionService struct {
	repo       repos.ParkingTransactionRepository
	obfuscator helpers.IDObfuscator
	validator  *validator.Validate
}

// NewParkingTransactionService creates a new instance of ParkingTransactionService
func NewParkingTransactionService(
	r repos.ParkingTransactionRepository,
	o helpers.IDObfuscator,
	v *validator.Validate,
) ParkingTransactionService {
	return &parkingTransactionService{
		repo:       r,
		obfuscator: o,
		validator:  v,
	}
}

// RegisterEntry handles the vehicle entry (gate in) logic
// It creates a new parking transaction with the following steps:
// 1. Validate the incoming request
// 2. Sanitize the license plate (uppercase, remove special characters)
// 3. Decode hashed IDs (VehicleTypeID, ZoneID)
// 4. Check for duplicate entry (vehicle already parked)
// 5. Validate zone capacity
// 6. Generate a unique ticket number
// 7. Create the transaction record with status "PARKED"
// 8. Return the transaction details with encoded IDs
func (s *parkingTransactionService) RegisterEntry(
	ctx context.Context,
	req requests.ParkingEntryRequest,
	operatorID uint,
) (responses.ParkingEntryResponse, error) {
	// Validate the request payload
	if err := s.validator.Struct(req); err != nil {
		return responses.ParkingEntryResponse{}, err
	}

	// Sanitize the license plate: uppercase and remove special characters
	sanitizedPlate := s.sanitizePlateNumber(req.PlateNumber)
	if sanitizedPlate == "" {
		return responses.ParkingEntryResponse{}, errors.New("invalid plate number format")
	}

	// Decode VehicleTypeID from hashed format
	vehicleTypeID, err := s.obfuscator.Decode(req.VehicleTypeID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("invalid vehicle type id")
	}

	// Decode ZoneID from hashed format
	zoneID, err := s.obfuscator.Decode(req.ZoneID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("invalid zone id")
	}

	// Check for duplicate entry: vehicle already parked
	// Query the database directly using GORM to find existing parked vehicle
	// This prevents duplicate entries for the same plate number
	var existingTransaction models.ParkingTransaction
	result := s.repo.GetDB().WithContext(ctx).
		Where("plate_number = ? AND status = ?", sanitizedPlate, "PARKED").
		First(&existingTransaction)
	
	if result.Error == nil {
		// Vehicle is already parked
		return responses.ParkingEntryResponse{}, errors.New("vehicle is already parked in the system")
	} else if result.Error.Error() != "record not found" {
		// Database error occurred
		return responses.ParkingEntryResponse{}, fmt.Errorf("failed to check duplicate entry: %w", result.Error)
	}

	// Validate zone capacity
	// Note: This requires fetching the zone to check current occupancy
	// For now, we'll add a TODO comment for future implementation
	// TODO: Implement zone occupancy check when ZoneService is available
	// if err := s.validateZoneCapacity(ctx, zoneID); err != nil {
	//     return responses.ParkingEntryResponse{}, err
	// }

	// Generate unique ticket number
	// Format: TCK-{timestamp}-{random alphanumeric}
	ticketNumber := s.generateTicketNumber()

	// Create the parking transaction model
	now := time.Now()
	transaction := models.ParkingTransaction{
		OperatorID:    operatorID,
		VehicleTypeID: vehicleTypeID,
		PlateNumber:   sanitizedPlate,
		RfidUID:       req.RfidUID,
		EntryTime:     now,
		Status:        "PARKED",
		PaymentStatus: "unpaid",
		Notes:         fmt.Sprintf("Ticket: %s", ticketNumber),
	}

	// Save the transaction to the database
	// Note: Create modifies the transaction in-place with the generated ID
	if err := s.repo.Create(ctx, &transaction); err != nil {
		return responses.ParkingEntryResponse{}, fmt.Errorf("failed to create parking transaction: %w", err)
	}

	// Encode the transaction ID for the response
	encodedTransactionID, err := s.obfuscator.Encode(transaction.ID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("failed to encode transaction id")
	}

	// Encode the vehicle type ID for the response
	encodedVehicleTypeID, err := s.obfuscator.Encode(vehicleTypeID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("failed to encode vehicle type id")
	}

	// Encode the zone ID for the response
	encodedZoneID, err := s.obfuscator.Encode(zoneID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("failed to encode zone id")
	}

	// Encode the operator ID for the response
	encodedOperatorID, err := s.obfuscator.Encode(operatorID)
	if err != nil {
		return responses.ParkingEntryResponse{}, errors.New("failed to encode operator id")
	}

	// Build and return the response
	return responses.ParkingEntryResponse{
		TransactionID:  encodedTransactionID,
		TicketNumber:   ticketNumber,
		EntryTime:      transaction.EntryTime,
		Status:         transaction.Status,
		VehicleTypeID:  encodedVehicleTypeID,
		ZoneID:         encodedZoneID,
		PlateNumber:    transaction.PlateNumber,
		OperatorID:     encodedOperatorID,
	}, nil
}

// generateTicketNumber generates a unique ticket number
// Format: TCK-{timestamp}-{random alphanumeric}
// Example: TCK-20260207050614-A7K9M2
func (s *parkingTransactionService) generateTicketNumber() string {
	timestamp := time.Now().Format("20060102150405")
	randomPart := s.generateRandomString(6)
	return fmt.Sprintf("TCK-%s-%s", timestamp, randomPart)
}

// generateRandomString generates a random alphanumeric string of the specified length
func (s *parkingTransactionService) generateRandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, nil)
		b[i] = charset[num.Uint64()%uint64(len(charset))]
	}
	return string(b)
}

// sanitizePlateNumber sanitizes the license plate number
// Converts to uppercase and removes special characters
// Allows only alphanumeric characters and hyphens
func (s *parkingTransactionService) sanitizePlateNumber(plate string) string {
	// Convert to uppercase
	plate = strings.ToUpper(strings.TrimSpace(plate))

	// Remove special characters except hyphens and spaces
	// Allow: A-Z, 0-9, hyphen (-), space
	re := regexp.MustCompile(`[^A-Z0-9\-\s]`)
	plate = re.ReplaceAllString(plate, "")

	// Remove extra spaces
	plate = strings.Join(strings.Fields(plate), " ")

	// Validate that the plate is not empty after sanitization
	if len(plate) == 0 {
		return ""
	}

	return plate
}
