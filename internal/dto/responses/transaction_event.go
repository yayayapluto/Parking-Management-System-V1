package responses

type TransactionEventResponse struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	EventType     string `json:"event_type"`
	OperatorID    string `json:"operator_id"`
	PhotoPath     string `json:"photo_path"`
	PlateDetected string `json:"plate_detected"`
	RfidDetected  string `json:"rfid_detected"`
	Notes         string `json:"notes"`
	CreatedAt     string `json:"created_at"`
}
