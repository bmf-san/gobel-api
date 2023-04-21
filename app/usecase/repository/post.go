package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Post is a repository interface for a post.
type Post interface {
	CountAllPublish() (int, error)
	CountAll() (int, error)
	CountAllPublishByKeyword(keyword string) (int, error)
	CountAllPublishByCategory(name string) (int, error)
	CountAllPublishByTag(name string) (int, error)
	FindAllPublish(page int, limit int) (domain.Posts, error)
	FindAllPublishByKeyword(page int, limit int, keyword string) (domain.Posts, error)
	FindAllPublishByCategory(page int, limit int, name string) (domain.Posts, error)
	FindAllPublishByTag(page int, limit int, name string) (domain.Posts, error)
	FindAll(page int, limit int) (domain.Posts, error)
	FindPublishByTitle(title string) (domain.Post, error)
	FindByID(id int) (domain.Post, error)
	Save(req request.StorePost) (int, error)
	SaveByID(req request.UpdatePost) error
	DeleteByID(id int) (int, error)
}
