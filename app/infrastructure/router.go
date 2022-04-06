package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/interfaces"
	"github.com/bmf-san/gobel-api/app/middleware"
	"github.com/bmf-san/goblin"
	"github.com/go-redis/redis/v7"
)

// Route sets the routing.
func Route(connMySQL *sql.DB, connRedis *redis.Client, l interfaces.Logger) *goblin.Router {
	ar := interfaces.AdminRepository{
		ConnMySQL: connMySQL,
		ConnRedis: connRedis,
	}
	jr := interfaces.JWTRepository{
		ConnRedis: connRedis,
	}

	mw := middleware.NewMiddleware(l, ar, jr)

	defaultController := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	authController := interfaces.NewAuthController(connMySQL, connRedis, l)
	postController := interfaces.NewPostController(connMySQL, connRedis, l)
	commentController := interfaces.NewCommentController(connMySQL, l)
	categoryController := interfaces.NewCategoryController(connMySQL, l)
	tagController := interfaces.NewTagController(connMySQL, l)

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
