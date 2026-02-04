package requests

type CreatePaymentMethodRequest struct {
	Code     string   `json:"code" validate:"required,max=20,alphanum"`
	Name     string   `json:"name" validate:"required,max=50"`
	Config   []string `json:"config"` // Kita terima sebagai slice string biasa
	IsActive bool     `json:"is_active"`
}

type UpdatePaymentMethodRequest struct {
	Name     string   `json:"name" validate:"required,max=50"`
	Config   []string `json:"config"`
	IsActive bool     `json:"is_active"`
}
