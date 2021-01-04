package response

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A IndexCategory represents the singular of IndexCategory.
type IndexCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A IndexCategories represents the plural of IndexCategory.
type IndexCategories []IndexCategory

// A IndexCategoryPrivate represents the singular of IndexCategoryPrivate.
type IndexCategoryPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A ShowCategory represents the singular of ShowCategory.
type ShowCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A ShowCategoryPrivate represents the singular of ShowCategoryPrivate.
type ShowCategoryPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A IndexCategoriesPrivate represents the plural of IndexCategoryPrivate.
type IndexCategoriesPrivate []IndexCategoryPrivate

// A StoreCategoryPrivate represents the singular of StoreCategoryPrivate.
type StoreCategoryPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MakeResponseIndex makes a response.
func MakeResponseIndexCategory(categories domain.Categories) []IndexCategory {
	var cat []IndexCategory
	for _, cs := range categories {
		c := IndexCategory{
			ID:   cs.ID,
			Name: cs.Name,
		}
		cat = append(cat, c)
	}
	return cat
}

// MakeResponseIndexPrivate makes a response.
func MakeResponseIndexCategoryPrivate(categories domain.Categories) []IndexCategoryPrivate {
	var cat []IndexCategoryPrivate
	for _, cs := range categories {
		c := IndexCategoryPrivate{
			ID:        cs.ID,
			Name:      cs.Name,
			CreatedAt: cs.CreatedAt,
			UpdatedAt: cs.UpdatedAt,
		}
		cat = append(cat, c)
	}
	return cat
}
