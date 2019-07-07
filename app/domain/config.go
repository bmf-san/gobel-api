package domain

import "time"

// A Configs represents the plural of config.
type Configs []Config

// A Config represents the singular of config.
type Config struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
