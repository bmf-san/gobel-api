package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
)

// A PostRepository is a repository interface for a post.
type PostRepository interface {
	CountAllPublish() (int, error)
	CountAll() (int, error)
	CountAllPublishByCategory(name string) (int, error)
	CountAllPublishByTag(name string) (int, error)
	FindAllPublish(page int, limit int) (domain.Posts, error)
	FindAllPublishByCategory(page int, limit int, name string) (domain.Posts, error)
	FindAllPublishByTag(page int, limit int, name string) (domain.Posts, error)
	FindAll(page int, limit int) (domain.Posts, error)
	FindPublishByTitle(title string) (domain.Post, error)
	FindByID(id int) (domain.Post, error)
	Save(req dto.RequestPost) (int, error)
	SaveByID(req dto.RequestPost, id int) error
	DeleteByID(id int) (int, error)
}
