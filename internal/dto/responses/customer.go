package responses

type CustomerResponse struct {
	ID                   string  `json:"id"`
	RfidUID              string  `json:"rfid_uid"`
	Name                 string  `json:"name"`
	Phone                string  `json:"phone"`
	RegistrationSourceID string  `json:"registration_source_id"`
	IsRegistered         bool    `json:"is_registered"`
	RegisteredAt         string  `json:"registered_at"`
	TotalVisits          int     `json:"total_visits"`
	TotalSpent           float64 `json:"total_spent"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
}
