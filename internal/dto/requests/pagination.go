package requests

type PaginationRequest struct {
	Page   int    `json:"page" query:"page" validate:"omitempty,min=1"`
	Limit  int    `json:"limit" query:"limit" validate:"omitempty,min=1,max=100"`
	Sort   string `json:"sort" query:"sort"`     // Contoh: "name asc"
	Search string `json:"search" query:"search"` // Global search (keyword)

	// Filter spesifik: Map[nama_kolom]nilai
	// Contoh: filters[is_active]=true&filters[vehicle_type]=CAR
	Filters map[string]interface{} `json:"filters" query:"filters"`
}

func (p *PaginationRequest) GetOffset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return (p.Page - 1) * p.GetLimit()
}

func (p *PaginationRequest) GetLimit() int {
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	return p.Limit
}
