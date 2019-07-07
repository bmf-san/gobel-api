package usecases

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A CategoryResponse is a success response.
type CategoryResponse struct{}

// MakeResponseHandleIndex makes a response.
func (r *CategoryResponse) MakeResponseHandleIndex(categories domain.Categories) ResponseCategories {
	var res []ResponseCategory
	for _, c := range categories {
		responseCategory := ResponseCategory{
			ID:   c.ID,
			Name: c.Name,
		}

		res = append(res, responseCategory)
	}

	return res
}

// MakeResponseHandleIndexPrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleIndexPrivate(categories domain.Categories) PrivateResponseCategories {
	var res []PrivateResponseCategory
	for _, c := range categories {
		responseCategory := PrivateResponseCategory{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}

		res = append(res, responseCategory)
	}

	return res
}

// MakeResponseHandleShow makes a response.
func (r *CategoryResponse) MakeResponseHandleShow(c domain.Category) ResponseCategory {
	return ResponseCategory{
		ID:   c.ID,
		Name: c.Name,
	}
}

// MakeResponseHandleShowPrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleShowPrivate(c domain.Category) PrivateResponseCategory {
	return PrivateResponseCategory{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// A ResponseCategory represents the singular of category for response.
type ResponseCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A ResponseCategories represents the plural of post for response.
type ResponseCategories []ResponseCategory

// A PrivateResponseCategory represents the singular of category for response.
type PrivateResponseCategory struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A PrivateResponseCategories represents the plural of post for response.
type PrivateResponseCategories []PrivateResponseCategory
