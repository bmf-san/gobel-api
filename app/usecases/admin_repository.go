package usecases

import "github.com/bmf-san/gobel-api/app/domain"

// An AdminRepository is a repository interface for an authentication.
type AdminRepository interface {
	FindByJWTAuth(req RequestJWTAuthHandleJWTAuth) (domain.Admin, error)
}
