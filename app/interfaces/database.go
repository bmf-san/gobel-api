package interfaces

import "database/sql"

// A DB represents a database connection.
type DB interface {
	GetConn() (*sql.DB, error)
}
