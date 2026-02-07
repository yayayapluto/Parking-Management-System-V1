package responses

type ReportCacheResponse struct {
	ID         string `json:"id"`
	ReportType string `json:"report_type"`
	CacheKey   string `json:"cache_key"`
	CacheData  string `json:"cache_data"`
	ExpiresAt  string `json:"expires_at"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
