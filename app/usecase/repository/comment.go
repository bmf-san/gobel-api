package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Comment is a repository interface for a comment.
type Comment interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Comments, error)
	FindByID(id int) (domain.Comment, error)
	Save(req request.StoreComment) (int, error)
	SaveStatusByID(req request.UpdateCommentStatus) error
}
