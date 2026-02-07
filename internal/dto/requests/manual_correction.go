package requests

type CreateManualCorrectionRequest struct {
	TransactionID  string `json:"transaction_id" validate:"required"`
	FieldCorrected string `json:"field_corrected" validate:"required,max=100"`
	OldValue       string `json:"old_value" validate:"required,max=255"`
	NewValue       string `json:"new_value" validate:"required,max=255"`
	CorrectedBy    string `json:"corrected_by" validate:"required"`
	Reason         string `json:"reason" validate:"required,max=500"`
}
