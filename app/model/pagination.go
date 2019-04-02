package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Pagination 分页结果
type Pagination struct {
	HasPrev  bool        `json:"has_prev"`
	HasNext  bool        `json:"has_next"`
	Pages    int         `json:"pages"`
	PrevPage int         `json:"prev_page"`
	NextPage int         `json:"next_page"`
	Items    interface{} `json:"items"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	Total    int         `json:"total"`
}

// Paginate 对查询query进行分页
func Paginate(query *gorm.DB, page, perPage int, items interface{}) (*Pagination, error) {
	var total int
	done := make(chan error, 1)
	go getTotal(query, &total, done)

	if err := query.Offset((page - 1) * perPage).Limit(perPage).Find(items).Error; err != nil {
		return nil, fmt.Errorf("pagination error: %v", err)
	}

	if err := <-done; err != nil {
		return nil, fmt.Errorf("pagination error: %v", err)
	}

	pages := (total + perPage - 1) / perPage
	hasPrev := page > 1
	hasNext := page < pages
	prevPage := page
	if hasPrev {
		prevPage = page - 1
	}
	nextPage := page
	if hasNext {
		nextPage = page + 1
	}
	return &Pagination{
		HasPrev:  hasPrev,
		HasNext:  hasNext,
		Pages:    pages,
		PrevPage: prevPage,
		NextPage: nextPage,
		Items:    items,
		Page:     page,
		PerPage:  perPage,
		Total:    total}, nil
}

func getTotal(query *gorm.DB, total *int, done chan error) {
	done <- query.Count(total).Error
}
