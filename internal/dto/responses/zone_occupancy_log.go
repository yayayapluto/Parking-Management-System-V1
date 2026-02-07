package responses

type ZoneOccupancyLogResponse struct {
	ID             string `json:"id"`
	ZoneID         string `json:"zone_id"`
	TransactionID  string `json:"transaction_id"`
	OperatorID     string `json:"operator_id"`
	OccupiedCount  int    `json:"occupied_count"`
	AvailableSlots int    `json:"available_slots"`
	EventType      string `json:"event_type"`
	Notes          string `json:"notes"`
	CreatedAt      string `json:"created_at"`
}
