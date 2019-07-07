package domain

import "time"

// A Admins represents the plural of admin.
type Admins []Admin

// A Admin represents the singular of admin.
type Admin struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
