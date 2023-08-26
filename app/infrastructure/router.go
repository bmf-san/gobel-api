package infrastructure

import (
	"database/sql"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/bmf-san/gobel-api/app/interfaces/controller"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/goblin"
	"github.com/go-redis/redis/v9"
)

// Route sets the routing.
func Route(connm *sql.DB, connr *redis.Client, l *slog.Logger) *goblin.Router {
	ar := repository.AdminRepository{
		ConnMySQL: connm,
		ConnRedis: connr,
	}
	jr := repository.JWT{
		ConnRedis: connr,
	}

	mw := NewMiddleware(l, ar, jr)

	defaultOPTIONSHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		w.Header().Set("Access-Control-Max-Age", "86400")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, Access-Control-Allow-Origin")
		w.WriteHeader(http.StatusNoContent)
	})
	authController := controller.NewAuthController(connm, connr, l)
	postController := controller.NewPostController(connm, connr, l)
	commentController := controller.NewCommentController(connm, l)
	categoryController := controller.NewCategoryController(connm, l)
	tagController := controller.NewTagController(connm, l)

	r := goblin.NewRouter()

	r.DefaultOPTIONSHandler = defaultOPTIONSHandler

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/healthcheck`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))

	r.UseGlobal(mw.CORS)

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/heap", pprof.Handler("heap"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler("/debug/pprof/block", pprof.Handler("block"))

	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts`, postController.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/search`, postController.IndexByKeyword())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/categories/:name`, postController.IndexByCategory())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/tags/:name`, postController.IndexByTag())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/posts/:title`, postController.Show())
	r.Methods(http.MethodPost).Use(mw.Recovery).Handler(`/posts/:title/comments`, commentController.Store())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/categories`, categoryController.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/categories/:name`, categoryController.Show())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/tags`, tagController.Index())
	r.Methods(http.MethodGet).Use(mw.Recovery).Handler(`/tags/:name`, tagController.Show())
	r.Methods(http.MethodPost).Use(mw.Recovery).Handler(`/signin`, authController.SignIn())
	r.Methods(http.MethodPost).Use(mw.Recovery, mw.Auth).Handler(`/private/signout`, authController.SignOut())
	r.Methods(http.MethodPost).Use(mw.Recovery, mw.Refresh).Handler(`/private/refresh`, authController.Refresh())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/me`, authController.ShowUserInfo())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/posts`, postController.IndexPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/posts/:id`, postController.ShowPrivate())
	r.Methods(http.MethodPost).Use(mw.Recovery, mw.Auth).Handler(`/private/posts`, postController.StorePrivate())
	r.Methods(http.MethodPatch).Use(mw.Recovery, mw.Auth).Handler(`/private/posts/:id`, postController.UpdatePrivate())
	r.Methods(http.MethodDelete).Use(mw.Recovery, mw.Auth).Handler(`/private/posts/:id`, postController.DestroyPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/comments`, commentController.IndexPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/comments/:id`, commentController.ShowPrivate())
	r.Methods(http.MethodPatch).Use(mw.Recovery, mw.Auth).Handler(`/private/comments/:id/status`, commentController.UpdateStatusPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/categories`, categoryController.IndexPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/categories/:id`, categoryController.ShowPrivate())
	r.Methods(http.MethodPost).Use(mw.Recovery, mw.Auth).Handler(`/private/categories`, categoryController.StorePrivate())
	r.Methods(http.MethodPatch).Use(mw.Recovery, mw.Auth).Handler(`/private/categories/:id`, categoryController.UpdatePrivate())
	r.Methods(http.MethodDelete).Use(mw.Recovery, mw.Auth).Handler(`/private/categories/:id`, categoryController.DestroyPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/tags`, tagController.IndexPrivate())
	r.Methods(http.MethodGet).Use(mw.Recovery, mw.Auth).Handler(`/private/tags/:id`, tagController.ShowPrivate())
	r.Methods(http.MethodPost).Use(mw.Recovery, mw.Auth).Handler(`/private/tags`, tagController.StorePrivate())
	r.Methods(http.MethodPatch).Use(mw.Recovery, mw.Auth).Handler(`/private/tags/:id`, tagController.UpdatePrivate())
	r.Methods(http.MethodDelete).Use(mw.Recovery, mw.Auth).Handler(`/private/tags/:id`, tagController.DestroyPrivate())

	return r
}
