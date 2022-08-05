package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
)

// A CategoryRepository is a repository interface for a comment.
type CategoryRepository interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Categories, error)
	FindByID(id int) (domain.Category, error)
	FindByName(name string) (domain.Category, error)
	Save(req dto.RequestCategory) (int, error)
	SaveByID(req dto.RequestCategory, id int) error
	DeleteByID(id int) (int, error)
}
