package usecases

import "github.com/bmf-san/gobel-api/app/domain"

// A CategoryRepository is a repository interface for a comment.
type CategoryRepository interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Categories, error)
	FindByID(id int) (domain.Category, error)
	FindByName(name string) (domain.Category, error)
	Save(req RequestCategory) (int, error)
	SaveByID(req RequestCategory, id int) error
	DeleteByID(id int) (int, error)
}
