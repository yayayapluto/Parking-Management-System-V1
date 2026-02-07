package responses

type RoleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PermissionID string `json:"permission_id"`
	Permission  *PermissionResponse `json:"permission,omitempty"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
