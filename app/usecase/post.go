package usecase

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A Post represents a Post.
type Post interface {
	Index(request.IndexPost) (domain.Posts, interactor.Pagination, *interactor.HTTPError)
	IndexByKeyword(request.IndexPostByKeyword) (domain.Posts, interactor.Pagination, *interactor.HTTPError)
	IndexByCategory(request.IndexPostByName) (domain.Posts, interactor.Pagination, *interactor.HTTPError)
	IndexByTag(request.IndexPostByName) (domain.Posts, interactor.Pagination, *interactor.HTTPError)
	IndexPrivate(request.IndexPost) (domain.Posts, interactor.Pagination, *interactor.HTTPError)
	Show(request.ShowPostByTitle) (domain.Post, *interactor.HTTPError)
	ShowPrivate(request.ShowPostByID) (domain.Post, *interactor.HTTPError)
	StorePrivate(request.StorePost) (domain.Post, *interactor.HTTPError)
	UpdatePrivate(request.UpdatePost) (domain.Post, *interactor.HTTPError)
	DestroyPrivate(request.DestroyPostByID) *interactor.HTTPError
}
