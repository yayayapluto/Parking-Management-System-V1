package requests

type CreateReportCacheRequest struct {
	ReportType string `json:"report_type" validate:"required,max=50"`
	CacheKey   string `json:"cache_key" validate:"required,max=255"`
	CacheData  string `json:"cache_data" validate:"required"`
	ExpiresAt  string `json:"expires_at" validate:"required"`
}

type UpdateReportCacheRequest struct {
	ReportType string `json:"report_type" validate:"omitempty,max=50"`
	CacheKey   string `json:"cache_key" validate:"omitempty,max=255"`
	CacheData  string `json:"cache_data" validate:"omitempty"`
	ExpiresAt  string `json:"expires_at" validate:"omitempty"`
}
