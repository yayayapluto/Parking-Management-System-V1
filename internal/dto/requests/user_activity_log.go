package requests

type CreateUserActivityLogRequest struct {
	UserID      string   `json:"user_id" validate:"required"`
	Action      string   `json:"action" validate:"required,max=50"`
	Module      string   `json:"module" validate:"required,max=50"`         // Ganti dari Resource
	Description string   `json:"description" validate:"omitempty,max=5000"` // Ganti dari Details
	IPAddress   string   `json:"ip_address" validate:"omitempty,ip"`
	UserAgent   string   `json:"user_agent" validate:"omitempty,max=500"`
	RequestData []string `json:"request_data" validate:"omitempty"` // Tambahkan ini
}
