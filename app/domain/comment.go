package domain

import "time"

// A Comments represents the plural of comment.
type Comments []Comment

// A Comment represents the singular of comment.
type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
