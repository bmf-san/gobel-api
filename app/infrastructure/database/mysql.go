package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/bmf-san/gobel-api/app/interfaces/database"
)

// A MySQLConn represents a mysql connection
type MySQLConn struct{}

// NewMySQLConn creates a MySQLConn struct.
func NewMySQLConn() database.MySQLConn {
	return &MySQLConn{}
}

// Conn get a mysql connection.
func (mc *MySQLConn) Conn() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open(os.Getenv("DB_DRIVER"), dataSourceName)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Second * 10)

	return conn, nil
}
