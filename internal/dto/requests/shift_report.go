package requests

import "time"

type CreateShiftReportRequest struct {
	UserID            string    `json:"user_id" validate:"required"`
	ShiftStart        time.Time `json:"shift_start" validate:"required"`
	ShiftEnd          time.Time `json:"shift_end" validate:"required"`
	OpeningCash       float64   `json:"opening_cash" validate:"min=0"`
	ClosingCash       float64   `json:"closing_cash" validate:"min=0"`
	TotalTransactions int       `json:"total_transactions" validate:"min=0"`
	TotalCash         float64   `json:"total_cash" validate:"min=0"`
	TotalQRIS         float64   `json:"total_qris" validate:"min=0"`
	Notes             string    `json:"notes" validate:"omitempty,max=5000"`
}

type UpdateShiftReportRequest struct {
	ClosingCash       *float64 `json:"closing_cash" validate:"omitempty,min=0"`
	TotalTransactions *int     `json:"total_transactions" validate:"omitempty,min=0"`
	TotalCash         *float64 `json:"total_cash" validate:"omitempty,min=0"`
	TotalQRIS         *float64 `json:"total_qris" validate:"omitempty,min=0"`
	Notes             string   `json:"notes" validate:"omitempty,max=5000"`
}
