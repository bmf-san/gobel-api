package response

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A IndexComment represents the singular of IndexComment.
type IndexComment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A IndexComments represents the plural of IndexComment.
type IndexComments []IndexComment

// A IndexCommentPrivate represents the singular of IndexCommentPrivate.
type IndexCommentPrivate struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A IndexCommentsPrivate represents the plural of IndexCommentsPrivate.
type IndexCommentsPrivate []IndexCommentPrivate

// A ShowCommentPrivate represents the singular of ShowCommentPrivate.
type ShowCommentPrivate struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A StoreAndUpdateComment represents the singular of StoreAndUpdateComment.
type StoreAndUpdateComment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MakeResponseIndexCommentPrivate makes a response.
func MakeResponseIndexCommentPrivate(comments domain.Comments) []IndexCommentPrivate {
	var cmt []IndexCommentPrivate
	for _, c := range comments {
		rc := IndexCommentPrivate{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}
		cmt = append(cmt, rc)
	}
	return cmt
}
