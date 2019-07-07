package domain

import "time"

// A Posts represents the plural of post.
type Posts []Post

// A Post represents the singular of post.
type Post struct {
	ID        int       `json:"id"`
	Admin     Admin     `json:"admin"`
	Category  Category  `json:"category"`
	Tags      Tags      `json:"tags"`
	Title     string    `json:"title"`
	MDBody    string    `json:"md_body"`
	HTMLBody  string    `json:"html_body"`
	Status    string    `json:"status"`
	Comments  Comments  `json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
