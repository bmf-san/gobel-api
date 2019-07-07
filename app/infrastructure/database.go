package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bmf-san/gobel-api/app/interfaces"
)

// A DB represents a database connection
type DB struct{}

// NewDB creates a DB struct.
func NewDB() interfaces.DB {
	return &DB{}
}

// GetConn get a database connection.
func (d *DB) GetConn() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
