package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
)

// A CommentRepository is a repository interface for a comment.
type CommentRepository interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Comments, error)
	FindByID(id int) (domain.Comment, error)
	Save(req dto.RequestComment) (int, error)
	SaveStatusByID(req dto.RequestCommentStatus, id int) error
}
