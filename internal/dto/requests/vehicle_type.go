package requests

type CreateVehicleTypeRequest struct {
	Code        string `json:"code" validate:"required,alphanum,max=10"`
	Name        string `json:"name" validate:"required,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type UpdateVehicleTypeRequest struct {
	Code        string `json:"code" validate:"omitempty,alphanum,max=10"`
	Name        string `json:"name" validate:"omitempty,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
	IsActive    *bool  `json:"is_active" validate:"omitempty"`
}
