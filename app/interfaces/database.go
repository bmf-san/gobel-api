package interfaces

import (
	"database/sql"

	"github.com/go-redis/redis/v7"
)

// A DB represents a database connection.
type DB interface {
	GetConnMySQL() (*sql.DB, error)
	GetConnRedis() (*redis.Client, error)
}
