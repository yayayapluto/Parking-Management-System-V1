package responses

type ZoneResponse struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	ZoneTypeID         string `json:"zone_type_id"`
	Location           string `json:"location"`
	MaximumCapacity    int    `json:"maximum_capacity"`
	DefaultFee         float64 `json:"default_fee"`
	Is24Hours          bool   `json:"is_24_hours"`
	OperatingHourStart string `json:"operating_hour_start"`
	OperatingHourEnd   string `json:"operating_hour_end"`
	IsActive           bool   `json:"is_active"`
	IsInMaintenance    bool   `json:"is_in_maintenance"`
	Description        string `json:"description"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}
