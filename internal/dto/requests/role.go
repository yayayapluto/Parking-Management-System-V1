package requests

type CreateRoleRequest struct {
	Name         string `json:"name" validate:"required,max=50"`
	PermissionID string `json:"permission_id" validate:"required"`
	CreatedBy    string `json:"created_by" validate:"omitempty,max=100"`
}

type UpdateRoleRequest struct {
	Name         string `json:"name" validate:"omitempty,max=50"`
	PermissionID string `json:"permission_id" validate:"omitempty"`
	UpdatedBy    string `json:"updated_by" validate:"omitempty,max=100"`
}
