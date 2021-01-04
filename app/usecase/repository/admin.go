package repository

import (
	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// An AdminRepository is a repository interface for an authentication.
type AdminRepository interface {
	FindByID(id int) (domain.Admin, error)
	FindByCredentials(req request.SignIn) (domain.Admin, error)
}
