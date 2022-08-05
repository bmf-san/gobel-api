package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/controller"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/goblin"
	"github.com/go-redis/redis/v9"
)

// Route sets the routing.
func Route(connm *sql.DB, connr *redis.Client, l domain.Logger) *goblin.Router {
	ar := repository.AdminRepository{
		ConnMySQL: connm,
		ConnRedis: connr,
	}
	jr := repository.JWTRepository{
		ConnRedis: connr,
	}

	mw := NewMiddleware(l, ar, jr)

	defaultController := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	authController := controller.NewAuthController(connm, connr, l)
	postController := controller.NewPostController(connm, connr, l)
	commentController := controller.NewCommentController(connm, l)
	categoryController := controller.NewCategoryController(connm, l)
	tagController := controller.NewTagController(connm, l)

	r := goblin.NewRouter()

	r.Methods(http.MethodGet).Handler(`/healthcheck`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/posts`, postController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/posts`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/posts/categories/:name`, postController.IndexByCategory())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/posts/categories/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/posts/tags/:name`, postController.IndexByTag())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/posts/tags/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/posts/:title`, postController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/posts/:title`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain).Handler(`/posts/:title/comments`, commentController.Store())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/posts/:title/comments`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/categories`, categoryController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/categories`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/categories/:name`, categoryController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/categories/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/tags`, tagController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/tags`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain).Handler(`/tags/:name`, tagController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/tags/:name`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain).Handler(`/signin`, authController.SignIn())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/signin`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain, mw.Auth).Handler(`/private/signout`, authController.SignOut())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/signout`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain, mw.Refresh).Handler(`/private/refresh`, authController.Refresh())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/refresh`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/me`, authController.ShowMe())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/me`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/posts`, postController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/posts`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/posts/:id`, postController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain, mw.Auth).Handler(`/private/posts`, postController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/posts`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORSMain, mw.Auth).Handler(`/private/posts/:id`, postController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORSMain, mw.Auth).Handler(`/private/posts/:id`, postController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/comments`, commentController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/comments`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/comments/:id`, commentController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/comments/:id`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORSMain, mw.Auth).Handler(`/private/comments/:id/status`, commentController.UpdateStatusPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/comments/:id/status`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/categories`, categoryController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/categories`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/categories/:id`, categoryController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain, mw.Auth).Handler(`/private/categories`, categoryController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/categories`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORSMain, mw.Auth).Handler(`/private/categories/:id`, categoryController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORSMain, mw.Auth).Handler(`/private/categories/:id`, categoryController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/tags`, tagController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/tags`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORSMain, mw.Auth).Handler(`/private/tags/:id`, tagController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/tags/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORSMain, mw.Auth).Handler(`/private/tags`, tagController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/tags`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORSMain, mw.Auth).Handler(`/private/tags/:id`, tagController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/tags/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORSMain, mw.Auth).Handler(`/private/tags/:id`, tagController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORSPreflight).Handler(`/private/tags/:id`, defaultController)

	return r
}
