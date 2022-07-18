package dto

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A TagResponse is a success response.
type TagResponse struct{}

// MakeResponseHandleIndex makes a response.
func (r *TagResponse) MakeResponseHandleIndex(tags domain.Tags) (int, []byte, error) {
	var rt []ResponseTag

	for _, t := range tags {
		responseTag := ResponseTag{
			ID:   t.ID,
			Name: t.Name,
		}
		rt = append(rt, responseTag)
	}

	res, err := json.Marshal(rt)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleIndexPrivate makes a response.
func (r *TagResponse) MakeResponseHandleIndexPrivate(tags domain.Tags) (int, []byte, error) {
	var rt []PrivateResponseTag

	for _, t := range tags {
		responseTag := PrivateResponseTag{
			ID:        t.ID,
			Name:      t.Name,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}
		rt = append(rt, responseTag)
	}

	res, err := json.Marshal(rt)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleShow makes a response.
func (r *TagResponse) MakeResponseHandleShow(t domain.Tag) (int, []byte, error) {
	res, err := json.Marshal(ResponseTag{
		ID:   t.ID,
		Name: t.Name,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleShowPrivate makes a response.
func (r *TagResponse) MakeResponseHandleShowPrivate(t domain.Tag) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseTag{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleStorePrivate makes a response.
func (r *TagResponse) MakeResponseHandleStorePrivate(t domain.Tag) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseTag{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, res, nil
}

// MakeResponseHandleUpdatePrivate makes a response.
func (r *TagResponse) MakeResponseHandleUpdatePrivate(t domain.Tag) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseTag{
		ID:        t.ID,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// A ResponseTag represents the singular of tag for response.
type ResponseTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A ResponseTags represents the plural of tag for response.
type ResponseTags []ResponseTag

// A PrivateResponseTag represents the singular of tag for response.
type PrivateResponseTag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A PrivateResponseTags represents the plural of tag for response.
type PrivateResponseTags []PrivateResponseTag
