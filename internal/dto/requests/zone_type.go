package requests

type CreateZoneTypeRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

type UpdateZoneTypeRequest struct {
	Name        string `json:"name" validate:"omitempty,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
}
