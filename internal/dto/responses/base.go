package responses

import "parking-management-system-v1/internal/dto/requests"

// BaseResponse buat single object atau pesan sukses doang
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageResponse buat yang ada paginationnya
type PageResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Meta        `json:"meta"`
}

// Meta buat info halaman
type Meta struct {
	CurrentPage int                    `json:"current_page"`
	TotalPage   int                    `json:"total_page"`
	TotalData   int64                  `json:"total_data"`
	Limit       int                    `json:"limit"`
	Sort        string                 `json:"sort,omitempty"`
	Search      string                 `json:"search,omitempty"`
	Filters     map[string]interface{} `json:"filters,omitempty"` // Balikin filter yang aktif
}

func CreateMeta(p requests.PaginationRequest, totalData int64) Meta {
	limit := p.GetLimit()
	totalPage := int(totalData) / limit
	if int(totalData)%limit > 0 {
		totalPage++
	}

	return Meta{
		CurrentPage: p.Page,
		TotalPage:   totalPage,
		TotalData:   totalData,
		Limit:       limit,
		Sort:        p.Sort,
		Search:      p.Search,
		Filters:     p.Filters,
	}
}
