package responses

type ZoneRateResponse struct {
	ID                string  `json:"id"`
	ZoneID            string  `json:"zone_id"`
	VehicleTypeID     string  `json:"vehicle_type_id"`
	HourlyRate        float64 `json:"hourly_rate"`
	DailyMaxRate      float64 `json:"daily_max_rate"`
	FreeMinutes       int     `json:"free_minutes"`
	IsWeekend         bool    `json:"is_weekend"`
	IsHoliday         bool    `json:"is_holiday"`
	HolidayID         *string `json:"holiday_id"`
	ValidFrom         string  `json:"valid_from"`
	ValidTo           string  `json:"valid_to"`
	EffectiveHourFrom string  `json:"effective_hour_from"`
	EffectiveHourTo   string  `json:"effective_hour_to"`
	IsActive          bool    `json:"is_active"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}
