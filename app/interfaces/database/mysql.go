package database

import (
	"database/sql"
)

// A MySQLConn represents a database connection.
type MySQLConn interface {
	Conn() (*sql.DB, error)
}
