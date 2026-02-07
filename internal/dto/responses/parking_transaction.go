package responses

import "time"

// ParkingEntryResponse represents the response payload for successful vehicle entry
type ParkingEntryResponse struct {
	// TransactionID is the hashed ID of the created parking transaction
	TransactionID string `json:"transaction_id"`

	// TicketNumber is the unique ticket number generated for this entry (e.g., TCK-20260207050614-ABC123)
	TicketNumber string `json:"ticket_number"`

	// EntryTime is the timestamp when the vehicle entered the parking zone
	EntryTime time.Time `json:"entry_time"`

	// Status is the current status of the transaction (should be "PARKED" for new entries)
	Status string `json:"status"`

	// VehicleTypeID is the hashed ID of the vehicle type
	VehicleTypeID string `json:"vehicle_type_id"`

	// ZoneID is the hashed ID of the parking zone
	ZoneID string `json:"zone_id"`

	// PlateNumber is the vehicle's license plate number
	PlateNumber string `json:"plate_number"`

	// OperatorID is the hashed ID of the operator who recorded the entry
	OperatorID string `json:"operator_id"`
}

// ParkingExitResponse represents the response payload for successful vehicle exit
type ParkingExitResponse struct {
	// TransactionID is the hashed ID of the parking transaction
	TransactionID string `json:"transaction_id"`

	// RfidUID is the RFID tag UID of the vehicle
	RfidUID string `json:"rfid_uid"`

	// PlateNumber is the vehicle's license plate number
	PlateNumber string `json:"plate_number"`

	// EntryTime is the timestamp when the vehicle entered the parking zone
	EntryTime time.Time `json:"entry_time"`

	// ExitTime is the timestamp when the vehicle exited the parking zone
	ExitTime time.Time `json:"exit_time"`

	// TotalDuration is the total parking duration in minutes
	TotalDuration int `json:"total_duration"`

	// TotalFee is the calculated parking fee
	TotalFee float64 `json:"total_fee"`

	// PaymentStatus is the status of the payment (paid, unpaid, etc.)
	PaymentStatus string `json:"payment_status"`

	// PaymentMethod is the method used for payment (manual or qris)
	PaymentMethod string `json:"payment_method"`
}
