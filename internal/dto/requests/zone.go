package requests

type CreateZoneRequest struct {
	Name               string  `json:"name" validate:"required,max=100"`
	ZoneTypeID         string  `json:"zone_type_id" validate:"required"`
	Location           string  `json:"location" validate:"required,max=255"`
	MaximumCapacity    int     `json:"maximum_capacity" validate:"required,min=1"`
	DefaultFee         float64 `json:"default_fee" validate:"required,min=0"`
	Is24Hours          bool    `json:"is_24_hours" validate:"omitempty"`
	OperatingHourStart string  `json:"operating_hour_start" validate:"omitempty"`
	OperatingHourEnd   string  `json:"operating_hour_end" validate:"omitempty"`
	Description        string  `json:"description" validate:"omitempty,max=1000"`
}

type UpdateZoneRequest struct {
	Name               string  `json:"name" validate:"omitempty,max=100"`
	ZoneTypeID         string  `json:"zone_type_id" validate:"omitempty"`
	Location           string  `json:"location" validate:"omitempty,max=255"`
	MaximumCapacity    *int    `json:"maximum_capacity" validate:"omitempty,min=1"`
	DefaultFee         *float64 `json:"default_fee" validate:"omitempty,min=0"`
	Is24Hours          *bool   `json:"is_24_hours" validate:"omitempty"`
	OperatingHourStart string  `json:"operating_hour_start" validate:"omitempty"`
	OperatingHourEnd   string  `json:"operating_hour_end" validate:"omitempty"`
	IsActive           *bool   `json:"is_active" validate:"omitempty"`
	IsInMaintenance    *bool   `json:"is_in_maintenance" validate:"omitempty"`
	Description        string  `json:"description" validate:"omitempty,max=1000"`
}
