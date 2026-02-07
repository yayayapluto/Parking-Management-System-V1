package requests

import "time"

type CreateDailySettlementRequest struct {
	Date              time.Time `json:"date" validate:"required"`
	TotalTransactions int       `json:"total_transactions" validate:"required,min=0"`
	TotalRevenue      float64   `json:"total_revenue" validate:"required,min=0"`
	TotalCash         float64   `json:"total_cash" validate:"required,min=0"`
	TotalQRIS         float64   `json:"total_qris" validate:"required,min=0"`
	TotalRefunds      float64   `json:"total_refunds" validate:"required,min=0"`
	TotalLostTickets  int       `json:"total_lost_tickets" validate:"required,min=0"`
}

type ReconcileDailySettlementRequest struct {
	ReconciledBy string `json:"reconciled_by" validate:"required"`
	Notes        string `json:"notes" validate:"omitempty,max=500"`
}
