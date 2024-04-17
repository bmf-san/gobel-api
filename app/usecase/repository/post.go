package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Post is a repository interface for a post.
type Post interface {
	CountAllPublic() (int, error)
	CountAll() (int, error)
	CountAllPublicByKeyword(keyword string) (int, error)
	CountAllPublicByCategory(name string) (int, error)
	CountAllPublicByTag(name string) (int, error)
	FindAllPublic(page int, limit int) (domain.Posts, error)
	FindAllPublicByKeyword(page int, limit int, keyword string) (domain.Posts, error)
	FindAllPublicByCategory(page int, limit int, name string) (domain.Posts, error)
	FindAllPublicByTag(page int, limit int, name string) (domain.Posts, error)
	FindAll(page int, limit int) (domain.Posts, error)
	FindPublicByTitle(title string) (domain.Post, error)
	FindByID(id int) (domain.Post, error)
	Save(req request.StorePost) (int, error)
	SaveByID(req request.UpdatePost) error
	DeleteByID(id int) (int, error)
}
