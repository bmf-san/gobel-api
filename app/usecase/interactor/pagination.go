package interactor

import (
	"fmt"
	"math"
	"strconv"
)

const (
	defaultPage  = 1
	defaultLimit = 1
)

// A Pagination represents the singular of Pagination.
type Pagination struct {
	Count     string `json:"count"`
	PageCount string `json:"page_count"`
	Page      string `json:"page"`
	Limit     string `json:"limit"`
}

// NewPagination creates a Pagination.
func (p Pagination) NewPagination(count int, page int, limit int) Pagination {
	if page == 0 {
		page = defaultPage
	}
	if limit == 0 {
		limit = defaultLimit
	}
	return Pagination{
		Count:     strconv.Itoa(count),
		PageCount: fmt.Sprint(math.Ceil(float64(count) / float64(limit))),
		Page:      strconv.Itoa(page),
		Limit:     strconv.Itoa(limit),
	}
}
