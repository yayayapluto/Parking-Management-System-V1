package responses

type ShiftReportResponse struct {
	ID                string  `json:"id"`
	UserID            string  `json:"user_id"`
	ShiftStart        string  `json:"shift_start"`
	ShiftEnd          string  `json:"shift_end"`
	OpeningCash       float64 `json:"opening_cash"`
	ClosingCash       float64 `json:"closing_cash"`
	TotalTransactions int     `json:"total_transactions"`
	TotalCash         float64 `json:"total_cash"`
	TotalQRIS         float64 `json:"total_qris"`
	TotalRevenue      float64 `json:"total_revenue"`
	Discrepancy       float64 `json:"discrepancy"`
	Notes             string  `json:"notes"`
	CreatedAt         string  `json:"created_at"`
}
