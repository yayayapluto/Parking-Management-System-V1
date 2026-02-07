package requests

type CreateLostTicketFeeRequest struct {
	TransactionID string  `json:"transaction_id" validate:"required"`
	OriginalFee   float64 `json:"original_fee" validate:"required,min=0"`
	LostTicketFee float64 `json:"lost_ticket_fee" validate:"required,min=0"`
	ProcessedBy   string  `json:"processed_by" validate:"required"`
	Reason        string  `json:"reason" validate:"omitempty,max=255"`
}
