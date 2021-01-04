package repository

import "github.com/bmf-san/gobel-api/app/domain"

// A JWT is a repository interface for jwt.
type JWT interface {
	FindIDByAccessUUID(au string) (int, error)
	FindIDByRefreshUUID(au string) (int, error)
	SaveID(id int) (domain.JWT, error)
	DeleteByAccessUUID(au string) (int64, error)
	DeleteByRefreshUUID(au string) (int64, error)
}
