package usecases

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A CommentResponse is a success response.
type CommentResponse struct{}

// MakeResponseHandleIndexPrivate makes a response.
func (r *CommentResponse) MakeResponseHandleIndexPrivate(comments domain.Comments) PrivateResponseComments {
	var res []PrivateResponseComment
	for _, c := range comments {
		rc := PrivateResponseComment{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
		}
		res = append(res, rc)
	}

	return res
}

// MakeResponseHandleShowPrivate makes a response.
func (r *CommentResponse) MakeResponseHandleShowPrivate(c domain.Comment) PrivateResponseComment {
	return PrivateResponseComment{
		ID:        c.ID,
		PostID:    c.PostID,
		Body:      c.Body,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
	}
}

// A ResponseComment represents the singular of comment for response.
type ResponseComment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// A ResponseComments represents the plural of comment for response.
type ResponseComments []ResponseComment

// A PrivateResponseComment represents the singular of comment for response.
type PrivateResponseComment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

// A PrivateResponseComments represents the plural of comment for response.
type PrivateResponseComments []PrivateResponseComment
