package repository

import "github.com/bmf-san/gobel-api/app/domain"

// A JWTRepository is a repository interface for jwt.
type JWTRepository interface {
	FindIDByAccessUUID(au string) (int, error)
	FindIDByRefreshUUID(au string) (int, error)
	SaveID(id int) (domain.JWT, error)
	DeleteByAccessUUID(au string) (int64, error)
	DeleteByRefreshUUID(au string) (int64, error)
}
