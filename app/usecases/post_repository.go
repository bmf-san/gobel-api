package usecases

import (
	"github.com/bmf-san/gobel-api/app/domain"
)

// A PostRepository is a repository interface for a post.
type PostRepository interface {
	CountAllPublish() (int, error)
	CountAll() (int, error)
	FindAllPublish(page int, limit int) (domain.Posts, error)
	FindAllPublishByCategory(page int, limit int, name string) (domain.Posts, error)
	FindAllPublishByTag(page int, limit int, name string) (domain.Posts, error)
	FindAll(page int, limit int) (domain.Posts, error)
	FindByTitle(title string) (domain.Post, error)
	FindByID(id int) (domain.Post, error)
	Save(req RequestPost) error
	SaveByID(req RequestPost, id int) error
	DeleteByID(id int) (int, error)
}
