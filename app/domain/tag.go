package domain

import "time"

// A Tags represents the plural of tag.
type Tags []Tag

// A Tag represetns the singular of tag.
type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
