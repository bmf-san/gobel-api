package domain

import "time"

// A Categories represents the plural of category.
type Categories []Category

// A Category represents the singular of category.
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
