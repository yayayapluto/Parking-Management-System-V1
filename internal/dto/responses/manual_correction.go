package responses

type ManualCorrectionResponse struct {
	ID             string `json:"id"`
	TransactionID  string `json:"transaction_id"`
	FieldCorrected string `json:"field_corrected"`
	OldValue       string `json:"old_value"`
	NewValue       string `json:"new_value"`
	Reason         string `json:"reason"`
	CorrectedBy    string `json:"corrected_by"`
	CorrectedAt    string `json:"corrected_at"`
	CreatedAt      string `json:"created_at"`
}
