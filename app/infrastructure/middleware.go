package infrastructure

import (
	"log/slog"
	"net/http"
	"os"
	"runtime"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/controller"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// Middleware represents the plural of middelware.
type Middleware struct {
	logger          domain.Logger
	adminRepository repository.AdminRepository
	JWT             repository.JWT
}

// NewLogger creates a Middleware.
func NewMiddleware(l *Logger, ar repository.AdminRepository, jr repository.JWT) *Middleware {
	return &Middleware{
		logger:          l,
		adminRepository: ar,
		JWT:             jr,
	}
}

// Log is a middleware for logging. It logs the access log. It also adds a trace id to the context.
func (mw *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := mw.logger.WithTraceID(r.Context())

		mw.logger.InfoContext(ctx, "access log", slog.String("http_method", r.Method), slog.String("path", r.URL.Path), slog.String("remote_addr", r.RemoteAddr), slog.String("user_agent", r.UserAgent()))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Recovery is a middleware for recovering from panic.
func (mw *Middleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case string:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e)
				case runtime.Error:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.Error())
				case error:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.Error())
				default:
					mw.logger.ErrorContext(r.Context(), "[panic] "+e.(string))
				}
				controller.JSONResponse(w, http.StatusInternalServerError, nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// Auth is a middleware for authentication.
func (mw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var j domain.JWT

		verifiedToken, err := j.GetVerifiedAccessToken(r.Header.Get("Authorization"))
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		accessUUID, err := j.GetAccessUUID(verifiedToken)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := mw.JWT.FindIDByAccessUUID(accessUUID)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = mw.adminRepository.FindByID(adminID)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
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
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		refreshUUID, err := j.GetRefreshUUID(verifiedToken)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		adminID, err := mw.JWT.FindIDByRefreshUUID(refreshUUID)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
			controller.JSONResponse(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = mw.adminRepository.FindByID(adminID)
		if err != nil {
			mw.logger.ErrorContext(r.Context(), err.Error())
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
