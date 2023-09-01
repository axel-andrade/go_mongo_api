package entities

type PaginationOptions struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Search string `json:"search"`
	Sort   string `json:"order"`
}

func (p *PaginationOptions) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *PaginationOptions) SetLimit(limit int) {
	if limit <= 0 {
		p.Limit = 10
	} else {
		p.Limit = limit
	}
}

func (p *PaginationOptions) SetPage(page int) {
	if page <= 0 {
		p.Page = 1
	} else {
		p.Page = page
	}
}

func (p *PaginationOptions) SetSearch(search string) {
	p.Search = search
}

func (p *PaginationOptions) SetSort(sort string) {
	if sort == "" {
		p.Sort = "id desc"
	} else {
		p.Sort = sort
	}
}

func BuildPaginationOptions(limit int, page int, sort string, search string) (*PaginationOptions, error) {
	var paginationOptions PaginationOptions

	paginationOptions.SetLimit(limit)
	paginationOptions.SetPage(page)
	paginationOptions.SetSearch(search)
	paginationOptions.SetSort(sort)

	return &paginationOptions, nil
}
