package requests

type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	CreatedBy   string `json:"created_by" validate:"omitempty,max=100"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name" validate:"omitempty,max=100"`
	Description string `json:"description" validate:"omitempty,max=1000"`
	UpdatedBy   string `json:"updated_by" validate:"omitempty,max=100"`
}
