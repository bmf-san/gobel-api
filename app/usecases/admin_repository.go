package usecases

import "github.com/bmf-san/gobel-api/app/domain"

// An AdminRepository is a repository interface for an authentication.
type AdminRepository interface {
	FindByID(id int) (domain.Admin, error)
	FindByCredentials(req RequestCredential) (domain.Admin, error)
}
