package usecase

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A Category represents a Category.
type Category interface {
	Index(request.IndexCategory) (domain.Categories, interactor.Pagination, *interactor.HTTPError)
	IndexPrivate(request.IndexCategory) (domain.Categories, interactor.Pagination, *interactor.HTTPError)
	Show(request.ShowCategoryByName) (domain.Category, *interactor.HTTPError)
	ShowPrivate(request.ShowCategoryByID) (domain.Category, *interactor.HTTPError)
	StorePrivate(request.StoreCategory) (domain.Category, *interactor.HTTPError)
	UpdatePrivate(request.UpdateCategory) (domain.Category, *interactor.HTTPError)
	DestroyPrivate(request.DestroyCategoryByID) *interactor.HTTPError
}
