package responses

type VehicleResponse struct {
	ID            string `json:"id"`
	CustomerID    string `json:"customer_id"`
	VehicleTypeID string `json:"vehicle_type_id"`
	PlateNumber   string `json:"plate_number"`
	Description   string `json:"description"`
	LastSeenAt    string `json:"last_seen_at"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
