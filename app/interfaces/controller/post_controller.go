package controller

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/dto"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
	"github.com/go-redis/redis/v9"
)

// A PostController is a controller for a post.
type PostController struct {
	PostInteractor interactor.PostInteractor
	Logger         domain.Logger
}

// NewPostController creates a PostController.
func NewPostController(connMySQL *sql.DB, connRedis *redis.Client, logger domain.Logger) *PostController {
	return &PostController{
		PostInteractor: interactor.PostInteractor{
			AdminRepository: &repository.AdminRepository{
				ConnMySQL: connMySQL,
				ConnRedis: connRedis,
			},
			PostRepository: &repository.PostRepository{
				ConnMySQL: connMySQL,
			},
			JWTRepository: &repository.JWTRepository{
				ConnRedis: connRedis,
			},
			JSONResponse: &dto.JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleIndex(w, r)
	})
}

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleIndexByCategory(w, r)
	})
}

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleIndexByTag(w, r)
	})
}

// IndexPrivate displays a listing of the resource.
func (pc *PostController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleIndexPrivate(w, r)
	})
}

// Show displays the specified resource.
func (pc *PostController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleShow(w, r)
	})
}

// ShowPrivate displays the specified resource.
func (pc *PostController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleShowPrivate(w, r)
	})
}

// StorePrivate stores a newly created resource in storage.
func (pc *PostController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleStorePrivate(w, r)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (pc *PostController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleUpdatePrivate(w, r)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (pc *PostController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.PostInteractor.HandleDestroyPrivate(w, r)
	})
}
