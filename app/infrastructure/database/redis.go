package database

import (
	"context"
	"fmt"
	"os"

	"github.com/bmf-san/gobel-api/app/interfaces/database"
	"github.com/go-redis/redis/v9"
)

// A RedisConn represents a redis connection
type RedisConn struct{}

// NewMySQLConn creates a RedisConn struct.
func NewRedisConn() database.RedisConn {
	return &RedisConn{}
}

// Conn get a redis connection.
func (rc *RedisConn) Conn() (*redis.Client, error) {
	dataSourceName := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn := redis.NewClient(&redis.Options{
		Addr:     dataSourceName,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	ctx := context.Background()
	_, err := conn.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
