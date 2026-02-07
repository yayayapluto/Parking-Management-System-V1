package requests

type CreateCustomerRequest struct {
	RfidUID              string `json:"rfid_uid" validate:"omitempty,max=50"`
	Name                 string `json:"name" validate:"required,max=100"`
	Phone                string `json:"phone" validate:"omitempty,max=20"`
	RegistrationSourceID string `json:"registration_source_id" validate:"omitempty"`
}

type UpdateCustomerRequest struct {
	RfidUID              string `json:"rfid_uid" validate:"omitempty,max=50"`
	Name                 string `json:"name" validate:"omitempty,max=100"`
	Phone                string `json:"phone" validate:"omitempty,max=20"`
	RegistrationSourceID string `json:"registration_source_id" validate:"omitempty"`
	IsRegistered         *bool  `json:"is_registered" validate:"omitempty"`
}
