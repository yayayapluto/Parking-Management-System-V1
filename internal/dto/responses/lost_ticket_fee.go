package responses

type LostTicketFeeResponse struct {
	ID            string  `json:"id"`
	TransactionID string  `json:"transaction_id"`
	OriginalFee   float64 `json:"original_fee"`
	LostTicketFee float64 `json:"lost_ticket_fee"`
	TotalCharged  float64 `json:"total_charged"`
	ProcessedBy   string  `json:"processed_by"`
	Reason        string  `json:"reason"`
	CreatedAt     string  `json:"created_at"`
}
