package responses

type DailySettlementResponse struct {
	ID                string  `json:"id"`
	Date              string  `json:"date"`
	TotalTransactions int     `json:"total_transactions"`
	TotalRevenue      float64 `json:"total_revenue"`
	TotalCash         float64 `json:"total_cash"`
	TotalQris         float64 `json:"total_qris"`
	TotalRefunds      float64 `json:"total_refunds"`
	TotalLostTickets  int     `json:"total_lost_tickets"`
	Variance          float64 `json:"variance"`
	ReconciledBy      string  `json:"reconciled_by"`
	ReconciledAt      string  `json:"reconciled_at"`
	CreatedAt         string  `json:"created_at"`
}
