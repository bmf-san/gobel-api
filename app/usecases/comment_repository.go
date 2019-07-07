package usecases

import "github.com/bmf-san/gobel-api/app/domain"

// A CommentRepository is a repository interface for a comment.
type CommentRepository interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Comments, error)
	FindByID(id int) (domain.Comment, error)
	Save(req RequestComment) error
	SaveStatusByID(req RequestCommentStatus, id int) error
}
