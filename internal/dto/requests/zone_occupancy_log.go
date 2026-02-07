package requests

type CreateZoneOccupancyLogRequest struct {
	ZoneID         string `json:"zone_id" validate:"required"`
	TransactionID  string `json:"transaction_id" validate:"required"`
	OperatorID     string `json:"operator_id" validate:"required"`
	OccupiedCount  int    `json:"occupied_count" validate:"required,min=0"`
	AvailableSlots int    `json:"available_slots" validate:"required,min=0"`
	EventType      string `json:"event_type" validate:"required,oneof=entry exit manual_adjustment system_sync"`
	Notes          string `json:"notes" validate:"omitempty,max=5000"`
}
