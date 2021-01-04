package usecase

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A Comment represents a Comment.
type Comment interface {
	IndexPrivate(request.IndexComment) (domain.Comments, interactor.Pagination, *interactor.HTTPError)
	ShowPrivate(request.ShowCommentByID) (domain.Comment, *interactor.HTTPError)
	Store(request.StoreComment) (domain.Comment, *interactor.HTTPError)
	UpdateStatusPrivate(request.UpdateCommentStatus) (domain.Comment, *interactor.HTTPError)
}
