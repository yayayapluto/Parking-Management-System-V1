package requests

type CreateSystemLogRequest struct {
	Level      int    `json:"level" validate:"required,min=0,max=5"`
	Service    string `json:"service" validate:"required,max=100"`
	Message    string `json:"message" validate:"required,max=1000"`
	Context    string `json:"context" validate:"omitempty"`
	StackTrace string `json:"stack_trace" validate:"omitempty"`
}
