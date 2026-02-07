package responses

type PermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Module      string `json:"module"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
