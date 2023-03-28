package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Category is a repository interface for a comment.
type Category interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Categories, error)
	FindByID(id int) (domain.Category, error)
	FindByName(name string) (domain.Category, error)
	Save(req request.StoreCategory) (int, error)
	SaveByID(req request.UpdateCategory) error
	DeleteByID(id int) (int, error)
}
