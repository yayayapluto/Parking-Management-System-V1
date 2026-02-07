package responses

type UserActivityLogResponse struct {
	ID          string   `json:"id"`
	UserID      string   `json:"user_id"`
	Action      string   `json:"action"`
	Module      string   `json:"module"`
	Description string   `json:"description"`
	IPAddress   string   `json:"ip_address"`
	UserAgent   string   `json:"user_agent"`
	RequestData []string `json:"request_data"`
	CreatedAt   string   `json:"created_at"`
}
