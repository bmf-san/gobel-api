package response

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A IndexTag represents the singular of IndexTag.
type IndexTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A IndexTags represents the plural of IndexTag.
type IndexTags []IndexTag

// A IndexTagPrivate represents the singular of IndexTagPrivate.
type IndexTagPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A ShowTag represents the singular of ShowTag.
type ShowTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A ShowTagPrivate represents the singular of ShowTagPrivate.
type ShowTagPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A IndexTagsPrivate represents the plural of IndexTagPrivate.
type IndexTagsPrivate []IndexTagPrivate

// A StoreTagPrivate represents the singular of StoreTagPrivate.
type StoreTagPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MakeResponseIndex makes a response.
func MakeResponseIndexTag(Tags domain.Tags) []IndexTag {
	var it []IndexTag
	for _, ts := range Tags {
		t := IndexTag{
			ID:   ts.ID,
			Name: ts.Name,
		}
		it = append(it, t)
	}
	return it
}

// MakeResponseIndexPrivate makes a response.
func MakeResponseIndexTagPrivate(Tags domain.Tags) []IndexTagPrivate {
	var its []IndexTagPrivate
	for _, ts := range Tags {
		t := IndexTagPrivate{
			ID:        ts.ID,
			Name:      ts.Name,
			CreatedAt: ts.CreatedAt,
			UpdatedAt: ts.UpdatedAt,
		}
		its = append(its, t)
	}
	return its
}
