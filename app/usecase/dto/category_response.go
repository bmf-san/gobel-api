package dto

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A CategoryResponse is a success response.
type CategoryResponse struct{}

// MakeResponseHandleIndex makes a response.
func (r *CategoryResponse) MakeResponseHandleIndex(categories domain.Categories) (int, []byte, error) {
	var cat []ResponseCategory
	for _, c := range categories {
		responseCategory := ResponseCategory{
			ID:   c.ID,
			Name: c.Name,
		}

		cat = append(cat, responseCategory)
	}

	res, err := json.Marshal(cat)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleIndexPrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleIndexPrivate(categories domain.Categories) (int, []byte, error) {
	var cat []PrivateResponseCategory
	for _, c := range categories {
		responseCategory := PrivateResponseCategory{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}

		cat = append(cat, responseCategory)
	}

	res, err := json.Marshal(cat)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleShow makes a response.
func (r *CategoryResponse) MakeResponseHandleShow(c domain.Category) (int, []byte, error) {
	res, err := json.Marshal(ResponseCategory{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleShowPrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleShowPrivate(c domain.Category) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseCategory{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleStorePrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleStorePrivate(c domain.Category) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseCategory{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, res, nil
}

// MakeResponseHandleUpdatePrivate makes a response.
func (r *CategoryResponse) MakeResponseHandleUpdatePrivate(c domain.Category) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseCategory{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
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
