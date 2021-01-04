package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Tag is a repository interface for a post.
type Tag interface {
	CountAll() (int, error)
	FindAll(page int, limit int) (domain.Tags, error)
	FindByID(id int) (domain.Tag, error)
	FindByName(name string) (domain.Tag, error)
	Save(req request.StoreTag) (int, error)
	SaveByID(req request.UpdateTag) error
	DeleteByID(id int) (int, error)
}
