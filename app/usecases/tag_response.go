package usecases

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A TagResponse is a success response.
type TagResponse struct{}

// MakeResponseHandleIndex makes a response.
func (r *TagResponse) MakeResponseHandleIndex(tags domain.Tags) ResponseTags {
	var res []ResponseTag

	for _, t := range tags {
		responseTag := ResponseTag{
			ID:   t.ID,
			Name: t.Name,
		}
		res = append(res, responseTag)
	}

	return res
}

// MakeResponseHandleIndexPrivate makes a response.
func (r *TagResponse) MakeResponseHandleIndexPrivate(tags domain.Tags) PrivateResponseTags {
	var res []PrivateResponseTag

	for _, t := range tags {
		responseTag := PrivateResponseTag{
			ID:        t.ID,
			Name:      t.Name,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}
		res = append(res, responseTag)
	}

	return res
}

// MakeResponseHandleShow makes a response.
func (r *TagResponse) MakeResponseHandleShow(c domain.Tag) ResponseTag {
	return ResponseTag{
		ID:   c.ID,
		Name: c.Name,
	}
}

// MakeResponseHandleShowPrivate makes a response.
func (r *TagResponse) MakeResponseHandleShowPrivate(c domain.Tag) PrivateResponseTag {
	return PrivateResponseTag{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
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
