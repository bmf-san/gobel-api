package infrastructure

import (
	"net/http"
	"os"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/controller"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
)

// Middleware represents the plural of middelware.
type Middleware struct {
	logger          domain.Logger
	adminRepository repository.AdminRepository
	JWT             repository.JWT
}

// NewLogger creates a Middleware.
func NewMiddleware(l domain.Logger, ar repository.AdminRepository, jr repository.JWT) *Middleware {
	return &Middleware{
		logger:          l,
		adminRepository: ar,
		JWT:             jr,
	}
}

// Auth is a middleware for authentication.
func (mw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var j domain.JWT

		verifiedToken, err := j.GetVerifiedAccessToken(r.Header.Get("Authorization"))
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		accessUUID, err := j.GetAccessUUID(verifiedToken)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := mw.JWT.FindIDByAccessUUID(accessUUID)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = mw.adminRepository.FindByID(adminID)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Refresh is a middleware for refreshing a access token by refresh token.
func (mw *Middleware) Refresh(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var j domain.JWT

		verifiedToken, err := j.GetVerifiedRefreshToken(r.Header.Get("Authorization"))

		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		refreshUUID, err := j.GetRefreshUUID(verifiedToken)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := mw.JWT.FindIDByRefreshUUID(refreshUUID)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = mw.adminRepository.FindByID(adminID)
		if err != nil {
			mw.logger.Error(err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// CORSMain is a middleware for main requests.
func (mw *Middleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Pagination-Count, Pagination-Pagecount, Pagination-Page, Pagination-Limit")
		next.ServeHTTP(w, r)
	})
}

// CORSPreflight is a middleware for preflight requests.
func (mw *Middleware) CORSPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, Access-Control-Allow-Origin")
		next.ServeHTTP(w, r)
	})
}
