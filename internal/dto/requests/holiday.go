package requests

type CreateHolidayRequest struct {
	Date         string  `json:"date" validate:"required" example:"2026-01-01"`
	Name         string  `json:"name" validate:"required,max=100"`
	AffectsRates float64 `json:"affects_rates" validate:"min=0"`
}

type UpdateHolidayRequest struct {
	Name         string  `json:"name" validate:"max=100"`
	AffectsRates float64 `json:"affects_rates" validate:"min=0"`
}
