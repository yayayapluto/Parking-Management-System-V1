package requests

type CreateOCRLogRequest struct {
	TransactionID    string  `json:"transaction_id" validate:"required"`
	OCRType          string  `json:"ocr_type" validate:"required,oneof=entry exit verification"`
	ImagePath        string  `json:"image_path" validate:"required,max=500"`
	RawOCRResult     string  `json:"raw_ocr_result" validate:"omitempty"`
	DetectedText     string  `json:"detected_text" validate:"omitempty,max=255"`
	ValidatedPlate   string  `json:"validated_plate" validate:"omitempty,max=20"`
	ConfidenceScore  float64 `json:"confidence_score" validate:"min=0,max=100"`
	ProcessingTimeMs int     `json:"processing_time_ms" validate:"min=0"`
	Status           string  `json:"status" validate:"required,oneof=pending success failed timeout"`
	ErrorMessage     string  `json:"error_message" validate:"omitempty,max=500"`
}
