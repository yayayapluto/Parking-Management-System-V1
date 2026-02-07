package requests

type CreateVehicleRequest struct {
	CustomerID    string `json:"customer_id" validate:"omitempty"`
	VehicleTypeID string `json:"vehicle_type_id" validate:"required"`
	PlateNumber   string `json:"plate_number" validate:"required,max=20"`
	Description   string `json:"description" validate:"omitempty,max=1000"`
}

type UpdateVehicleRequest struct {
	CustomerID    string `json:"customer_id" validate:"omitempty"`
	VehicleTypeID string `json:"vehicle_type_id" validate:"omitempty"`
	PlateNumber   string `json:"plate_number" validate:"omitempty,max=20"`
	Description   string `json:"description" validate:"omitempty,max=1000"`
}
