package responses

import (
	"fmt"
	"net/url"
	"parking-management-system-v1/internal/dto/requests"
	"strconv"
)

// BaseResponse buat single object atau pesan sukses doang
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
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
	CurrentPage int   `json:"current_page"`
	TotalPage   int   `json:"total_page"`
	TotalData   int64 `json:"total_data"`
	Limit       int   `json:"limit"`
	// Navigasi Angka
	FirstPage int  `json:"first_page"`
	LastPage  int  `json:"last_page"`
	PrevPage  *int `json:"prev_page"` // Pakai pointer biar bisa null kalau gak ada
	NextPage  *int `json:"next_page"`
	// Navigasi URL
	CurrentUrl string  `json:"current_url"`
	FirstUrl   string  `json:"first_url"`
	LastUrl    string  `json:"last_url"`
	PrevUrl    *string `json:"prev_url"`
	NextUrl    *string `json:"next_url"`
	// Search & Filter
	Sort    string                 `json:"sort,omitempty"`
	Search  string                 `json:"search,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
}

func CreateMeta(p requests.PaginationRequest, totalData int64, baseUrl string) Meta {
	limit := p.GetLimit()
	totalPage := int((totalData + int64(limit) - 1) / int64(limit))
	if totalPage == 0 {
		totalPage = 1
	}

	// Helper buat generate URL dengan query params yang sama tapi beda page
	buildUrl := func(page int) string {
		u, _ := url.Parse(baseUrl)
		q := u.Query()
		q.Set("page", strconv.Itoa(page))
		q.Set("limit", strconv.Itoa(limit))
		if p.Search != "" {
			q.Set("search", p.Search)
		}
		if p.Sort != "" {
			q.Set("sort", p.Sort)
		}
		// Tambahkan filters jika ada
		for k, v := range p.Filters {
			q.Set(fmt.Sprintf("filters[%s]", k), fmt.Sprintf("%v", v))
		}
		u.RawQuery = q.Encode()
		return u.String()
	}

	meta := Meta{
		CurrentPage: p.Page,
		TotalPage:   totalPage,
		TotalData:   totalData,
		Limit:       limit,
		FirstPage:   1,
		LastPage:    totalPage,
		CurrentUrl:  buildUrl(p.Page),
		FirstUrl:    buildUrl(1),
		LastUrl:     buildUrl(totalPage),
		Sort:        p.Sort,
		Search:      p.Search,
		Filters:     p.Filters,
	}

	// Logic Prev Page
	if p.Page > 1 {
		prev := p.Page - 1
		meta.PrevPage = &prev
		prevUrl := buildUrl(prev)
		meta.PrevUrl = &prevUrl
	}

	// Logic Next Page
	if p.Page < totalPage {
		next := p.Page + 1
		meta.NextPage = &next
		nextUrl := buildUrl(next)
		meta.NextUrl = &nextUrl
	}

	return meta
}
