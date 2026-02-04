package responses

type HolidayResponse struct {
	ID           string  `json:"id"`
	Date         string  `json:"date"`
	Name         string  `json:"name"`
	AffectsRates float64 `json:"affects_rates"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
