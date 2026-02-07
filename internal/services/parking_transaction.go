package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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

	// RegisterExit handles the vehicle exit (gate out) logic with fee calculation
	RegisterExit(ctx context.Context, req requests.ParkingExitRequest) (responses.ParkingExitResponse, error)
}

type parkingTransactionService struct {
	repo          repos.ParkingTransactionRepository
	zoneRepo      repos.ZoneRepository
	zoneRateRepo  repos.ZoneRateRepository
	vehicleRepo   repos.VehicleRepository
	customerRepo  repos.CustomerRepository
	obfuscator    helpers.IDObfuscator
	validator     *validator.Validate
	feeCalculator FeeCalculator
}

// NewParkingTransactionService creates a new instance of ParkingTransactionService
func NewParkingTransactionService(
	r repos.ParkingTransactionRepository,
	zr repos.ZoneRepository,
	zrr repos.ZoneRateRepository,
	vr repos.VehicleRepository,
	cr repos.CustomerRepository,
	o helpers.IDObfuscator,
	v *validator.Validate,
	fc FeeCalculator,
) ParkingTransactionService {
	return &parkingTransactionService{
		repo:          r,
		zoneRepo:      zr,
		zoneRateRepo:  zrr,
		vehicleRepo:   vr,
		customerRepo:  cr,
		obfuscator:    o,
		validator:     v,
		feeCalculator: fc,
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

	// RFID Detection: Query the customers table with rfid_uid
	// If found, mark as Registered; if not, treat as Guest
	var customer *models.Customer
	customerResult := s.repo.GetDB().WithContext(ctx).
		Where("rfid_uid = ?", req.RfidUID).
		First(&customer)

	isRegistered := customerResult.Error == nil
	var customerID *uint
	if isRegistered {
		customerID = &customer.ID
	}

	// Double Entry Fix: Check for active parking using rfid_uid (NOT plate number)
	// where exit_time IS NULL
	var existingTransaction models.ParkingTransaction
	result := s.repo.GetDB().WithContext(ctx).
		Where("rfid_uid = ? AND exit_time IS NULL", req.RfidUID).
		First(&existingTransaction)

	if result.Error == nil {
		// Vehicle is already parked
		return responses.ParkingEntryResponse{}, errors.New("vehicle is already parked in the system")
	} else if result.Error.Error() != "record not found" {
		// Database error occurred
		return responses.ParkingEntryResponse{}, fmt.Errorf("failed to check duplicate entry: %w", result.Error)
	}

	// Capacity Validation: Fetch the Zone data
	// If OccupiedCount >= MaximumCapacity, return a "Zone is Full" error
	var zone models.Zone
	zoneResult := s.repo.GetDB().WithContext(ctx).First(&zone, zoneID)
	if zoneResult.Error != nil {
		return responses.ParkingEntryResponse{}, fmt.Errorf("failed to fetch zone: %w", zoneResult.Error)
	}

	if zone.OccupiedCount >= zone.MaximumCapacity {
		return responses.ParkingEntryResponse{}, errors.New("zone is full")
	}

	// Generate unique ticket number
	// Format: TCK-{timestamp}-{random alphanumeric}
	ticketNumber := s.generateTicketNumber()

	// Atomic Entry: Wrap everything in a transaction
	// Create the transaction and update zone occupancy atomically
	now := time.Now()
	var transaction models.ParkingTransaction

	err = s.repo.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create the parking transaction model
		transaction = models.ParkingTransaction{
			CustomerID:    customerID,
			OperatorID:    operatorID,
			VehicleTypeID: vehicleTypeID,
			PlateNumber:   sanitizedPlate,
			RfidUID:       req.RfidUID,
			EntryTime:     now,
			Status:        "PARKED",
			PaymentStatus: "unpaid",
			Notes:         fmt.Sprintf("Ticket: %s", ticketNumber),
		}

		// Create the transaction
		if err := tx.Create(&transaction).Error; err != nil {
			return fmt.Errorf("failed to create parking transaction: %w", err)
		}

		// Increment occupied_count and decrement available_slots using gorm.Expr
		if err := tx.Model(&models.Zone{}).
			Where("id = ?", zoneID).
			Updates(map[string]interface{}{
				"occupied_count":  gorm.Expr("occupied_count + ?", 1),
				"available_slots": gorm.Expr("available_slots - ?", 1),
			}).Error; err != nil {
			return fmt.Errorf("failed to update zone occupancy: %w", err)
		}

		return nil
	})

	if err != nil {
		return responses.ParkingEntryResponse{}, err
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
		TransactionID: encodedTransactionID,
		TicketNumber:  ticketNumber,
		EntryTime:     transaction.EntryTime,
		Status:        transaction.Status,
		VehicleTypeID: encodedVehicleTypeID,
		ZoneID:        encodedZoneID,
		PlateNumber:   transaction.PlateNumber,
		OperatorID:    encodedOperatorID,
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

// RegisterExit handles the vehicle exit (gate out) logic with fee calculation
// It performs the following steps:
// 1. Validate the incoming request
// 2. Find the active parking transaction using RFID UID
// 3. Calculate the parking duration and fee
// 4. Update the transaction with exit details in an atomic transaction
// 5. Update zone occupancy (decrease occupied_count, increase available_slots)
// 6. Update vehicle stats if the customer is registered
// 7. Return the exit details for receipt printing
func (s *parkingTransactionService) RegisterExit(
	ctx context.Context,
	req requests.ParkingExitRequest,
) (responses.ParkingExitResponse, error) {
	// Validate the request payload
	if err := s.validator.Struct(req); err != nil {
		return responses.ParkingExitResponse{}, err
	}

	// Find the active parking transaction using RFID UID
	// Criteria: rfid_uid = req.RfidUID AND exit_time IS NULL AND status = 'PARKED'
	var transaction models.ParkingTransaction
	result := s.repo.GetDB().WithContext(ctx).
		Where("rfid_uid = ? AND exit_time IS NULL AND status = ?", req.RfidUID, "PARKED").
		First(&transaction)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return responses.ParkingExitResponse{}, errors.New("no active parking found for the given RFID")
		}
		return responses.ParkingExitResponse{}, fmt.Errorf("failed to find parking transaction: %w", result.Error)
	}

	// Get the transaction zone information to determine which zone the vehicle is in
	var transactionZone models.TransactionZone
	zoneResult := s.repo.GetDB().WithContext(ctx).
		Where("transaction_id = ? AND exit_time IS NULL", transaction.ID).
		First(&transactionZone)

	if zoneResult.Error != nil {
		return responses.ParkingExitResponse{}, fmt.Errorf("failed to find transaction zone: %w", zoneResult.Error)
	}

	// Accurate Duration: Calculate duration based on transactionZone.EntryTime (not the gate entry time)
	now := time.Now()
	feeResult, err := s.feeCalculator.CalculateFeeWithDB(
		ctx,
		transactionZone.EntryTime,
		now,
		transactionZone.ActualZoneID,
		transaction.VehicleTypeID,
		s.repo.GetDB(),
	)
	if err != nil {
		return responses.ParkingExitResponse{}, fmt.Errorf("failed to calculate fee: %w", err)
	}

	// Atomic Exit: Wrap in a Transaction block
	err = s.repo.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Update ParkingTransaction (exit time, total fee, status 'completed')
		transaction.ExitTime = &now
		transaction.DurationMinutes = feeResult.DurationMinutes
		transaction.BaseFee = feeResult.BaseFee
		transaction.TotalFee = feeResult.TotalFee
		transaction.Status = "completed"
		transaction.PaymentStatus = "paid"

		if err := tx.Save(&transaction).Error; err != nil {
			return fmt.Errorf("failed to update parking transaction: %w", err)
		}

		// Update the transaction zone with exit time
		transactionZone.ExitTime = &now
		transactionZone.DurationMinutes = feeResult.DurationMinutes
		transactionZone.ZoneFee = feeResult.TotalFee

		if err := tx.Save(&transactionZone).Error; err != nil {
			return fmt.Errorf("failed to update transaction zone: %w", err)
		}

		// Fix Occupancy Logic: Decrement occupied_count and increment available_slots using gorm.Expr
		// DO NOT use MaximumCapacity - 1
		if err := tx.Model(&models.Zone{}).
			Where("id = ?", transactionZone.ActualZoneID).
			Updates(map[string]interface{}{
				"occupied_count":  gorm.Expr("occupied_count - ?", 1),
				"available_slots": gorm.Expr("available_slots + ?", 1),
				"last_exit_time":  now,
			}).Error; err != nil {
			return fmt.Errorf("failed to update zone occupancy: %w", err)
		}

		// Update Vehicle last_seen_at and Customer total_visits (+1) if the user is registered
		if transaction.CustomerID != nil {
			// Update vehicle last_seen_at
			if err := tx.Model(&models.Vehicle{}).
				Where("customer_id = ? AND plate_number = ?", transaction.CustomerID, transaction.PlateNumber).
				Update("last_seen_at", now).Error; err != nil {
				return fmt.Errorf("failed to update vehicle stats: %w", err)
			}

			// Update customer total visits
			if err := tx.Model(&models.Customer{}).
				Where("id = ?", transaction.CustomerID).
				Update("total_visits", gorm.Expr("total_visits + ?", 1)).Error; err != nil {
				return fmt.Errorf("failed to update customer visits: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return responses.ParkingExitResponse{}, err
	}

	// Encode the transaction ID for the response
	encodedTransactionID, err := s.obfuscator.Encode(transaction.ID)
	if err != nil {
		return responses.ParkingExitResponse{}, errors.New("failed to encode transaction id")
	}

	// Build and return the response
	return responses.ParkingExitResponse{
		TransactionID: encodedTransactionID,
		RfidUID:       transaction.RfidUID,
		PlateNumber:   transaction.PlateNumber,
		EntryTime:     transaction.EntryTime,
		ExitTime:      now,
		TotalDuration: feeResult.DurationMinutes,
		TotalFee:      feeResult.TotalFee,
		PaymentStatus: transaction.PaymentStatus,
		PaymentMethod: req.PaymentMethod,
	}, nil
}
