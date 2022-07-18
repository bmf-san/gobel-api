package database

import (
	"github.com/go-redis/redis/v9"
)

// A RedisConn represents a database connection.
type RedisConn interface {
	Conn() (*redis.Client, error)
}
