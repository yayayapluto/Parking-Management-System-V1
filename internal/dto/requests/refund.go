package requests

type CreateRefundRequest struct {
	PaymentID     string  `json:"payment_id" validate:"required"`
	TransactionID string  `json:"transaction_id" validate:"required"`
	Amount        float64 `json:"amount" validate:"required,min=0"`
	Reason        string  `json:"reason" validate:"required,max=255"`
	RefundedBy    string  `json:"refunded_by" validate:"required"`
}

type UpdateRefundRequest struct {
	Amount     *float64 `json:"amount" validate:"omitempty,min=0"`
	Reason     string   `json:"reason" validate:"omitempty,max=255"`
	Status     string   `json:"status" validate:"omitempty,oneof=requested pending_approval approved rejected completed"`
	ApprovedBy string   `json:"approved_by" validate:"omitempty"`
}

type ApproveRefundRequest struct {
	ApprovedBy string `json:"approved_by" validate:"required"`
}

type RejectRefundRequest struct {
	Reason string `json:"reason" validate:"required,max=255"`
}
