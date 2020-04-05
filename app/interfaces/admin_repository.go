package interfaces

import (
	"database/sql"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecases"
	"github.com/go-redis/redis/v7"
)

// An AdminRepository is a repository for an authentication.
type AdminRepository struct {
	ConnMySQL *sql.DB
	ConnRedis *redis.Client
}

// FindByID returns the entity identified by the given id.
func (ar *AdminRepository) FindByID(adminID int) (domain.Admin, error) {
	const query = `
		SELECT
			id,
			name,
			email,
			password
		FROM
			admins
		WHERE
			id = ?
	`
	row, err := ar.ConnMySQL.Query(query, adminID)

	defer row.Close()

	var admin domain.Admin

	if err != nil {
		return admin, nil
	}

	var id int
	var name string
	var password string
	var email string
	row.Next()

	if err = row.Scan(&id, &name, &email, &password); err != nil {
		return admin, err
	}

	return domain.Admin{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

// FindByCredentials saves login session by the credential.
func (ar *AdminRepository) FindByCredentials(req usecases.RequestCredential) (domain.Admin, error) {
	const query = `
		SELECT
			id,
			name,
			email,
			password
		FROM
			admins
		WHERE
			email = ?
	`
	row, err := ar.ConnMySQL.Query(query, req.Email)

	defer row.Close()

	var admin domain.Admin

	if err != nil {
		return admin, nil
	}

	var id int
	var name string
	var password string
	var email string
	row.Next()

	if err = row.Scan(&id, &name, &email, &password); err != nil {
		return admin, err
	}

	return domain.Admin{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
