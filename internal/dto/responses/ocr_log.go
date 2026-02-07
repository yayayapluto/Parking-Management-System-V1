package responses

type OCRLogResponse struct {
	ID               string  `json:"id"`
	TransactionID    string  `json:"transaction_id"`
	OCRType          string  `json:"ocr_type"`
	ImagePath        string  `json:"image_path"`
	RawOCRResult     string  `json:"raw_ocr_result"`
	DetectedText     string  `json:"detected_text"`
	ValidatedPlate   string  `json:"validated_plate"`
	ConfidenceScore  float64 `json:"confidence_score"`
	ProcessingTimeMs int     `json:"processing_time_ms"`
	Status           string  `json:"status"`
	ErrorMessage     string  `json:"error_message"`
	CreatedAt        string  `json:"created_at"`
}
