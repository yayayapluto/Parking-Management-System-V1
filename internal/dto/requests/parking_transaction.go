package requests

// ParkingEntryRequest represents the request payload for vehicle entry (gate in)
type ParkingEntryRequest struct {
	// VehicleTypeID is the hashed ID of the vehicle type (e.g., "motorcycle", "car", "truck")
	VehicleTypeID string `json:"vehicle_type_id" validate:"required"`

	// ZoneID is the hashed ID of the parking zone where the vehicle is entering
	ZoneID string `json:"zone_id" validate:"required"`

	// PlateNumber is the vehicle's license plate number
	PlateNumber string `json:"plate_number" validate:"required,max=20"`

	// RfidUID is the optional RFID tag UID for automatic vehicle identification
	RfidUID string `json:"rfid_uid" validate:"omitempty,max=50"`

	// OperatorID is the hashed ID of the operator (parking attendant) recording the entry
	// If not provided, it will be extracted from JWT context
	OperatorID string `json:"operator_id" validate:"omitempty"`
}
