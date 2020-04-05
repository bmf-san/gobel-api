package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

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

// VerifyPassword verfies hashed password and requested password.
func (a *Admin) VerifyPassword(hashedPassword []byte, reqPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, reqPassword); err != nil {
		return err
	}

	return nil
}
