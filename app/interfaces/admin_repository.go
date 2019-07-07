package interfaces

import (
	"database/sql"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecases"
)

// An AdminRepository is a repository for an authentication.
type AdminRepository struct {
	Conn *sql.DB
}

// FindByJWTAuth returns the entity identified by the given email.
func (ar *AdminRepository) FindByJWTAuth(req usecases.RequestJWTAuthHandleJWTAuth) (admin domain.Admin, err error) {
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
	row, err := ar.Conn.Query(query, req.Email)

	defer row.Close()

	if err != nil {
		return
	}

	var id int
	var name string
	var password string
	var email string
	row.Next()
	if err = row.Scan(&id, &name, &email, &password); err != nil {
		return
	}
	admin.ID = id
	admin.Name = name
	admin.Email = email
	admin.Password = password

	return
}
