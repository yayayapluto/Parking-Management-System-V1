package responses

type SystemLogResponse struct {
	ID         string `json:"id"`
	Level      int    `json:"level"`
	Service    string `json:"service"`
	Message    string `json:"message"`
	Context    string `json:"context"`
	StackTrace string `json:"stack_trace"`
	CreatedAt  string `json:"created_at"`
}
