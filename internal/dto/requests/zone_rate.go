package requests

type CreateZoneRateRequest struct {
	ZoneID             string  `json:"zone_id" validate:"required"`
	VehicleTypeID      string  `json:"vehicle_type_id" validate:"required"`
	HourlyRate         float64 `json:"hourly_rate" validate:"required,min=0"`
	DailyMaxRate       float64 `json:"daily_max_rate" validate:"required,min=0"`
	FreeMinutes        int     `json:"free_minutes" validate:"omitempty,min=0"`
	IsWeekend          bool    `json:"is_weekend" validate:"omitempty"`
	IsHoliday          bool    `json:"is_holiday" validate:"omitempty"`
	HolidayID          *string `json:"holiday_id" validate:"omitempty"`
	ValidFrom          string  `json:"valid_from" validate:"required"`
	ValidTo            string  `json:"valid_to" validate:"required"`
	EffectiveHourFrom  string  `json:"effective_hour_from" validate:"omitempty"`
	EffectiveHourTo    string  `json:"effective_hour_to" validate:"omitempty"`
}

type UpdateZoneRateRequest struct {
	ZoneID             string   `json:"zone_id" validate:"omitempty"`
	VehicleTypeID      string   `json:"vehicle_type_id" validate:"omitempty"`
	HourlyRate         *float64 `json:"hourly_rate" validate:"omitempty,min=0"`
	DailyMaxRate       *float64 `json:"daily_max_rate" validate:"omitempty,min=0"`
	FreeMinutes        *int     `json:"free_minutes" validate:"omitempty,min=0"`
	IsWeekend          *bool    `json:"is_weekend" validate:"omitempty"`
	IsHoliday          *bool    `json:"is_holiday" validate:"omitempty"`
	HolidayID          *string  `json:"holiday_id" validate:"omitempty"`
	ValidFrom          string   `json:"valid_from" validate:"omitempty"`
	ValidTo            string   `json:"valid_to" validate:"omitempty"`
	EffectiveHourFrom  string   `json:"effective_hour_from" validate:"omitempty"`
	EffectiveHourTo    string   `json:"effective_hour_to" validate:"omitempty"`
	IsActive           *bool    `json:"is_active" validate:"omitempty"`
}
