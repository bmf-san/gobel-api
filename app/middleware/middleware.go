package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// middelware represents the singular of middleware.
type middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares represents the plural of middleware.
type Middlewares struct {
	middlewares []middleware
}

// NewMiddlewares creates a middlewares.
func NewMiddlewares(mws ...middleware) Middlewares {
	return Middlewares{
		middlewares: append([]middleware(nil), mws...),
	}
}

// Then handles chaining middlewares.
func (mws Middlewares) Then(h http.HandlerFunc) http.HandlerFunc {
	for i := range mws.middlewares {
		h = mws.middlewares[len(mws.middlewares)-1-i](h)
	}

	return h
}

// Auth is a middleware  for authentication.
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unauthorized")
					}
					return []byte(os.Getenv("JWT_SECRET_KEY")), nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(err.Error()))
					return
				}
				if token.Valid {
					next.ServeHTTP(w, r)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
}

// CORS is a middleware for CORS.
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		r.Header.Set("Access-Control-Max-Age", "86400")
		r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		r.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		r.Header.Set("Access-Control-Expose-Headers", "Content-Length")
		r.Header.Set("Access-Control-Allow-Credentials", "true")

		next.ServeHTTP(w, r)
	}
}
