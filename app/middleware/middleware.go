// Package middleware is a middleware package.
// NOTE: I feel this middleware implementation is off the clean architecture.
package middleware

import (
	"net/http"
	"os"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces"
	"github.com/bmf-san/gobel-api/app/usecases"
)

// middelware represents the singular of middleware.
type middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares represents the plural of middleware.
type Middlewares struct {
	middlewares []middleware
}

// Asset represents the plural of middelware services.
type Asset struct {
	jwtRepository   interfaces.JWTRepository
	adminRepository interfaces.AdminRepository
	logger          usecases.Logger
}

// NewAsset creates a assets.
func NewAsset(jwtRepository interfaces.JWTRepository, adminRepository interfaces.AdminRepository, logger usecases.Logger) Asset {
	return Asset{
		jwtRepository:   jwtRepository,
		adminRepository: adminRepository,
		logger:          logger,
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

		verifiedToken, err := j.GetVerifiedAccessToken(r.Header.Get("Authorization"))
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		accessUUID, err := j.GetAccessUUID(verifiedToken)
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := a.jwtRepository.FindIDByAccessUUID(accessUUID)
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = a.adminRepository.FindByID(adminID)
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Refresh is a middleware for refreshing a access token by refresh token.
func (a *Asset) Refresh(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jr interfaces.JSONResponse
		var j domain.JWT

		verifiedToken, err := j.GetVerifiedRefreshToken(r.Header.Get("Authorization"))

		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		refreshUUID, err := j.GetRefreshUUID(verifiedToken)
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := a.jwtRepository.FindIDByRefreshUUID(refreshUUID)
		if err != nil {
			a.logger.Error(err.Error())
			jr.HTTPStatus(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = a.adminRepository.FindByID(adminID)
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
