package usecase

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A Tag represents a Tag.
type Tag interface {
	Index(request.IndexTag) (domain.Tags, interactor.Pagination, *interactor.HTTPError)
	IndexPrivate(request.IndexTag) (domain.Tags, interactor.Pagination, *interactor.HTTPError)
	Show(request.ShowTagByName) (domain.Tag, *interactor.HTTPError)
	ShowPrivate(request.ShowTagByID) (domain.Tag, *interactor.HTTPError)
	StorePrivate(request.StoreTag) (domain.Tag, *interactor.HTTPError)
	UpdatePrivate(request.UpdateTag) (domain.Tag, *interactor.HTTPError)
	DestroyPrivate(request.DestroyTagByID) *interactor.HTTPError
}
