package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/interfaces"
	"github.com/bmf-san/gobel-api/app/middleware"
	"github.com/bmf-san/gobel-api/app/usecase"
	"github.com/bmf-san/goblin"
	"github.com/go-redis/redis/v7"
)

// defaultHandler is a handler for default.
func defaultHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})
}

// Route sets the routing.
func Route(connMySQL *sql.DB, connRedis *redis.Client, l usecase.Logger) *goblin.Router {
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
		return
	})
	authController := interfaces.NewAuthController(connMySQL, connRedis, l)
	postController := interfaces.NewPostController(connMySQL, connRedis, l)
	commentController := interfaces.NewCommentController(connMySQL, l)
	categoryController := interfaces.NewCategoryController(connMySQL, l)
	tagController := interfaces.NewTagController(connMySQL, l)

	r := goblin.NewRouter()

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/posts`, postController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/posts`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/posts/categories/:name`, postController.IndexByCategory())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/posts/categories/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/posts/tags/:name`, postController.IndexByTag())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/posts/tags/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/posts/:title`, postController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/posts/:title`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS).Handler(`/posts/:title/comments`, commentController.Store())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/posts/:title/comments`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/categories`, categoryController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/categories`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/categories/:name`, categoryController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/categories/:name`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/tags`, tagController.Index())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/tags`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS).Handler(`/tags/:name`, tagController.Show())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/tags/:name`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS).Handler(`/signin`, authController.SignIn())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/signin`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS, mw.Auth).Handler(`/private/signout`, authController.SignOut())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/signout`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS, mw.Refresh).Handler(`/private/refresh`, authController.Refresh())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/refresh`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/me`, authController.ShowMe())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/me`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/posts`, postController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/posts`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/posts/:id`, postController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS, mw.Auth).Handler(`/private/posts`, postController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/posts`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORS, mw.Auth).Handler(`/private/posts/:id`, postController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORS, mw.Auth).Handler(`/private/posts/:id`, postController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/posts/:id`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/comments`, commentController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/comments`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/comments/:id`, commentController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/comments/:id`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORS, mw.Auth).Handler(`/private/comments/:id/status`, commentController.UpdateStatusPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/comments/:id/status`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/categories`, categoryController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/categories`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/categories/:id`, categoryController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS, mw.Auth).Handler(`/private/categories`, categoryController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/categories`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORS, mw.Auth).Handler(`/private/categories/:id`, categoryController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORS, mw.Auth).Handler(`/private/categories/:id`, categoryController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/categories/:id`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/tags`, tagController.IndexPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/tags`, defaultController)

	r.Methods(http.MethodGet).Use(mw.CORS, mw.Auth).Handler(`/private/tags/:id`, tagController.ShowPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/tags/:id`, defaultController)

	r.Methods(http.MethodPost).Use(mw.CORS, mw.Auth).Handler(`/private/tags`, tagController.StorePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/tags`, defaultController)

	r.Methods(http.MethodPatch).Use(mw.CORS, mw.Auth).Handler(`/private/tags/:id`, tagController.UpdatePrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/tags/:id`, defaultController)

	r.Methods(http.MethodDelete).Use(mw.CORS, mw.Auth).Handler(`/private/tags/:id`, tagController.DestroyPrivate())
	r.Methods(http.MethodOptions).Use(mw.CORS).Handler(`/private/tags/:id`, defaultController)

	return r
}
