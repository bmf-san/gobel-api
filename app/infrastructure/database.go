package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bmf-san/gobel-api/app/interfaces"
	"github.com/go-redis/redis/v7"
)

// A DB represents a database connection
type DB struct{}

// NewDB creates a DB struct.
func NewDB() interfaces.DB {
	return &DB{}
}

// GetConnMySQL get a database connection.
func (d *DB) GetConnMySQL() (*sql.DB, error) {
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

// GetConnRedis get a database connection.
func (d *DB) GetConnRedis() (*redis.Client, error) {
	dataSourceName := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn := redis.NewClient(&redis.Options{
		Addr:     dataSourceName,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}

	return conn, nil
}
