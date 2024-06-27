package meta

import (
	"strconv"
)

type Meta struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	TotalPages int `json:"total_pages"`
	Total      int `json:"total"`
}

func New(page, pageSize, total int, pagLimitDef string) (*Meta, error) {

	if pageSize <= 0 {
		var err error
		pageSize, err = strconv.Atoi(pagLimitDef)
		if err != nil {
			return nil, err
		}
	}

	totalPages := 0
	if total >= 0 {
		totalPages = (total + pageSize - 1) / pageSize
		if page > totalPages {
			page = totalPages
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		Total:      total,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *Meta) Limit() int {
	return p.PageSize
}
