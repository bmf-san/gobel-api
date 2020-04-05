// Package middleware is a middleware package.
// NOTE: I feel this middleware implementation is off the clean architecture.
package middleware

import (
	"net/http"
	"os"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces"
	"github.com/bmf-san/gobel-api/app/usecases"
	"github.com/go-redis/redis/v7"
)

// middelware represents the singular of middleware.
type middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares represents the plural of middleware.
type Middlewares struct {
	middlewares []middleware
}

// Asset represents the plural of middelware services.
type Asset struct {
	connRedis *redis.Client
	logger    usecases.Logger
}

// NewAsset creates a assets.
func NewAsset(connRedis *redis.Client, logger usecases.Logger) Asset {
	return Asset{
		connRedis: connRedis,
		logger:    logger,
	}
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

// Auth is a middleware for authentication.
func (a *Asset) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jr interfaces.JSONResponse
		var j domain.JWT
		token := j.ExtractToken(r.Header.Get("Authorization"))

		var err error
		_, err = j.VerifyToken(token, os.Getenv("JWT_ACCESS_TOKEN_SECRET"))
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// CORS is a middleware for CORS.
func (a *Asset) CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, Access-Control-Allow-Origin")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Pagination-Count, Pagination-Pagecount, Pagination-Page, Pagination-Limit")

		next.ServeHTTP(w, r)
	}
}
