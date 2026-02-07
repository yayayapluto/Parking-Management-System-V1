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
