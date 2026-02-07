package requests

// ParkingEntryRequest represents the request payload for vehicle entry (gate in)
type ParkingEntryRequest struct {
	// VehicleTypeID is the hashed ID of the vehicle type (e.g., "motorcycle", "car", "truck")
	VehicleTypeID string `json:"vehicle_type_id" validate:"required"`

	// ZoneID is the hashed ID of the parking zone where the vehicle is entering
	ZoneID string `json:"zone_id" validate:"required"`

	// PlateNumber is the vehicle's license plate number
	PlateNumber string `json:"plate_number" validate:"required,min=3,max=20"`

	// RfidUID is the RFID tag UID for automatic vehicle identification (required)
	RfidUID string `json:"rfid_uid" validate:"required,max=50"`

	// OperatorID is the hashed ID of the operator (parking attendant) recording the entry
	// If not provided, it will be extracted from JWT context
	OperatorID string `json:"operator_id" validate:"omitempty"`
}

// ParkingExitRequest represents the request payload for vehicle exit (gate out)
type ParkingExitRequest struct {
	// RfidUID is the RFID tag UID for identifying the vehicle
	RfidUID string `json:"rfid_uid" validate:"required,max=50"`

	// PaymentMethod is the method used for payment (manual or qris)
	PaymentMethod string `json:"payment_method" validate:"required,oneof=manual qris"`
}
