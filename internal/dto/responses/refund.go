package responses

type RefundResponse struct {
	ID            string  `json:"id"`
	PaymentID     string  `json:"payment_id"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
	Reason        string  `json:"reason"`
	RefundedBy    string  `json:"refunded_by"`
	ApprovedBy    string  `json:"approved_by"`
	Status        string  `json:"status"`
	RefundedAt    string  `json:"refunded_at"`
	CreatedAt     string  `json:"created_at"`
}
