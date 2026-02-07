package requests

type CreateTransactionEventRequest struct {
	TransactionID string `json:"transaction_id" validate:"required"`
	EventType     string `json:"event_type" validate:"required,oneof=entry exit manual_entry manual_exit correction lost_ticket"`
	OperatorID    string `json:"operator_id" validate:"required"`
	PhotoPath     string `json:"photo_path" validate:"omitempty,max=500"`
	PlateDetected string `json:"plate_detected" validate:"omitempty,max=20"`
	RfidDetected  string `json:"rfid_detected" validate:"omitempty,max=50"`
	Notes         string `json:"notes" validate:"omitempty,max=5000"`
}
