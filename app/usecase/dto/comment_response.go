package dto

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A CommentResponse is a success response.
type CommentResponse struct{}

// MakeResponseHandleIndexPrivate makes a response.
func (r *CommentResponse) MakeResponseHandleIndexPrivate(comments domain.Comments) (int, []byte, error) {
	var cmt []PrivateResponseComment
	for _, c := range comments {
		rc := PrivateResponseComment{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}
		cmt = append(cmt, rc)
	}

	res, err := json.Marshal(cmt)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleStore makes a response.
func (r *CommentResponse) MakeResponseHandleStore(c domain.Comment) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseComment{
		ID:        c.ID,
		PostID:    c.PostID,
		Body:      c.Body,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, res, nil
}

// MakeResponseHandleShowPrivate makes a response.
func (r *CommentResponse) MakeResponseHandleShowPrivate(c domain.Comment) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseComment{
		ID:        c.ID,
		PostID:    c.PostID,
		Body:      c.Body,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleUpdateStatusPrivate makes a response.
func (r *CommentResponse) MakeResponseHandleUpdateStatusPrivate(c domain.Comment) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseComment{
		ID:        c.ID,
		PostID:    c.PostID,
		Body:      c.Body,
		Status:    c.Status,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// A ResponseComment represents the singular of comment for response.
type ResponseComment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	UpdatedAt time.Time `json:"updated_at"`
}

// A PrivateResponseComments represents the plural of comment for response.
type PrivateResponseComments []PrivateResponseComment
