package usecase

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// An Auth represents a Auth.
type Auth interface {
	SignIn(request.SignIn) (domain.JWT, *interactor.HTTPError)
	SignOut(request.SignOut) *interactor.HTTPError
	Refresh(request.Refresh) (domain.JWT, *interactor.HTTPError)
	ShowUserInfo(request.ShowUserInfo) (domain.Admin, *interactor.HTTPError)
}
