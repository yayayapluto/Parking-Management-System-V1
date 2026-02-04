package requests

type CreateCustomerRegistSourceRequest struct {
	Name        string `json:"name" validate:"required,max=50"`
	Description string `json:"description" validate:"max=1000"`
}

type UpdateCustomerRegistSourceRequest struct {
	Name        string `json:"name" validate:"max=50"`
	Description string `json:"description" validate:"max=1000"`
}
