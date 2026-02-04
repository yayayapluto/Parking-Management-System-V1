package responses

type PaymentMethodResponse struct {
	ID        string   `json:"id"` // HashID
	Code      string   `json:"code"`
	Name      string   `json:"name"`
	Config    []string `json:"config"`
	IsActive  bool     `json:"is_active"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
