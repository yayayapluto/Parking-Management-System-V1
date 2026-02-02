package responses

type VehicleTypeResponse struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	IsActive    bool   `json:"is_active"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
