package domain

import (
	"html"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// A JWTAuth represents the singular of jwt authentication.
type JWTAuth struct {
	Token string `json:"token"`
}

// SignIn compares the password stored in the database with the password obtained from the request.
func (ja *JWTAuth) SignIn(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// NewJWT creates a jwt newly.
func (ja *JWTAuth) NewJWT(admin Admin) (token string, err error) {
	j := jwt.New(jwt.SigningMethodHS256)
	claims := j.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["id"] = admin.ID
	claims["email"] = admin.Email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	return j.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

// Extract extract values from a token.
func (ja *JWTAuth) Extract() (map[string]interface{}, error) {
	bearer := strings.Split(ja.Token, " ")
	token := html.EscapeString(bearer[1])
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
