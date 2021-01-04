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

// Route sets the routing.
func Route(connMySQL *sql.DB, connRedis *redis.Client, logger usecase.Logger) *goblin.Router {
	jwtRepository := interfaces.JWTRepository{
		ConnRedis: connRedis,
	}
	adminRepository := interfaces.AdminRepository{
		ConnMySQL: connMySQL,
		ConnRedis: connRedis,
	}

	asset := middleware.NewAsset(jwtRepository, adminRepository, logger)

	publicMws := middleware.NewMiddlewares(asset.CORS)
	privateMws := middleware.NewMiddlewares(asset.CORS, asset.Auth)
	refreshMws := middleware.NewMiddlewares(asset.CORS, asset.Refresh)

	authController := interfaces.NewAuthController(connMySQL, connRedis, logger)
	postController := interfaces.NewPostController(connMySQL, connRedis, logger)
	commentController := interfaces.NewCommentController(connMySQL, logger)
	categoryController := interfaces.NewCategoryController(connMySQL, logger)
	tagController := interfaces.NewTagController(connMySQL, logger)

	r := goblin.NewRouter()

	r.GET("/posts", publicMws.Then(postController.Index))
	r.GET("/posts/categories/:name", publicMws.Then(postController.IndexByCategory))
	r.GET("/posts/tags/:name", publicMws.Then(postController.IndexByTag))
	r.GET("/posts/:title", publicMws.Then(postController.Show))
	r.POST("/posts/:title/comments", publicMws.Then(commentController.Store))
	r.GET("/categories", publicMws.Then(categoryController.Index))
	r.GET("/categories/:name", publicMws.Then(categoryController.Show))
	r.GET("/tags", publicMws.Then(tagController.Index))
	r.GET("/tags/:name", publicMws.Then(tagController.Show))
	r.POST("/signin", publicMws.Then(authController.SignIn))
	r.OPTION("/signin", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.POST("/private/signout", privateMws.Then(authController.SignOut))
	r.OPTION("/private/signout", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.POST("/private/refresh", refreshMws.Then(authController.Refresh))
	r.OPTION("/private/refresh", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/me", privateMws.Then(authController.ShowMe))
	r.OPTION("/private/me", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/posts", privateMws.Then(postController.IndexPrivate))
	r.OPTION("/private/posts", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/posts/:id", privateMws.Then(postController.ShowPrivate))
	r.OPTION("/private/posts/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.POST("/private/posts", privateMws.Then(postController.StorePrivate))
	r.OPTION("/private/posts", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.PATCH("/private/posts/:id", privateMws.Then(postController.UpdatePrivate))
	r.OPTION("/private/posts/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.DELETE("/private/posts/:id", privateMws.Then(postController.DestroyPrivate))
	r.OPTION("/private/posts/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/comments", privateMws.Then(commentController.IndexPrivate))
	r.OPTION("/private/comments", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/comments/:id", privateMws.Then(commentController.ShowPrivate))
	r.OPTION("/private/comments/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.PATCH("/private/comments/:id/status", privateMws.Then(commentController.UpdateStatusPrivate))
	r.OPTION("/private/comments/:id/status", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/categories", privateMws.Then(categoryController.IndexPrivate))
	r.OPTION("/private/categories", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/categories/:id", privateMws.Then(categoryController.ShowPrivate))
	r.OPTION("/private/categories/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.POST("/private/categories", privateMws.Then(categoryController.StorePrivate))
	r.OPTION("/private/categories", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.PATCH("/private/categories/:id", privateMws.Then(categoryController.UpdatePrivate))
	r.OPTION("/private/categories/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.DELETE("/private/categories/:id", privateMws.Then(categoryController.DestroyPrivate))
	r.OPTION("/private/categories/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/tags", privateMws.Then(tagController.IndexPrivate))
	r.OPTION("/private/tags", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.GET("/private/tags/:id", privateMws.Then(tagController.ShowPrivate))
	r.OPTION("/private/tags/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.POST("/private/tags", privateMws.Then(tagController.StorePrivate))
	r.OPTION("/private/tags", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.PATCH("/private/tags/:id", privateMws.Then(tagController.UpdatePrivate))
	r.OPTION("/private/tags/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))
	r.DELETE("/private/tags/:id", privateMws.Then(tagController.DestroyPrivate))
	r.OPTION("/private/tags/:id", publicMws.Then(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})))

	return r
}
