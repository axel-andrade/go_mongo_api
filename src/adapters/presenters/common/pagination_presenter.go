package common_ptr

import (
	"go_mongo_api/src/entities"
	"math"
)

type PaginateResult struct {
	Total       uint64 `json:"total,omitempty" example:"50" description:"O número total de registros existentes"`
	Limit       uint64 `json:"limit,omitempty" example:"20" minimum:"1" maximum:"100" description:"O número máximo de documentos a serem retornados por página"`
	Page        uint64 `json:"page,omitempty" example:"1" minimum:"1" description:"O número da página atual"`
	TotalPages  int    `json:"total_pages,omitempty" example:"50" minimum:"1" description:"O número total de páginas"`
	HasPrevPage bool   `json:"has_prev_page" example:"true" description:"Indica se há uma página anterior disponível"`
	HasNextPage bool   `json:"has_next_page" example:"true" description:"Indica se há uma próxima página disponível"`
	PrevPage    uint64 `json:"prev_page,omitempty" example:"1" minimum:"1" description:"O número da página anterior"`
	NextPage    uint64 `json:"next_page,omitempty" example:"2" minimum:"1" description:"O número da próxima página"`
}

type PaginationPresenter struct{}

func BuildPaginationPresenter() *PaginationPresenter {
	return &PaginationPresenter{}
}

func (p *PaginationPresenter) Format(pagination entities.PaginationOptions, totalDocs uint64) PaginateResult {
	totalPages := int(math.Ceil(float64(totalDocs) / float64(pagination.Limit)))

	var result PaginateResult

	result.Total = totalDocs
	result.TotalPages = totalPages
	result.Limit = uint64(pagination.Limit)
	result.Page = uint64(pagination.Page)
	result.HasPrevPage = int64(pagination.Page) > 1
	result.HasNextPage = int64(pagination.Page) < int64(totalPages)

	if result.HasPrevPage {
		result.PrevPage = uint64(pagination.Page) - 1
	}

	if result.HasNextPage {
		result.NextPage = uint64(pagination.Page) + 1
	}

	return result
}
