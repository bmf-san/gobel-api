package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecase"
	"github.com/go-redis/redis/v7"
)

// A PostController is a controller for a post.
type PostController struct {
	PostInteractor usecase.PostInteractor
	Logger         usecase.Logger
}

// NewPostController creates a PostController.
func NewPostController(connMySQL *sql.DB, connRedis *redis.Client, logger usecase.Logger) *PostController {
	return &PostController{
		PostInteractor: usecase.PostInteractor{
			AdminRepository: &AdminRepository{
				ConnMySQL: connMySQL,
				ConnRedis: connRedis,
			},
			PostRepository: &PostRepository{
				ConnMySQL: connMySQL,
			},
			JWTRepository: &JWTRepository{
				ConnRedis: connRedis,
			},
			JSONResponse: &JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleIndex(w, r)
	return
}

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleIndexByCategory(w, r)
	return
}

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleIndexByTag(w, r)
	return
}

// IndexPrivate displays a listing of the resource.
func (pc *PostController) IndexPrivate(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleIndexPrivate(w, r)
	return
}

// Show displays the specified resource.
func (pc *PostController) Show(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleShow(w, r)
	return
}

// ShowPrivate displays the specified resource.
func (pc *PostController) ShowPrivate(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleShowPrivate(w, r)
	return
}

// StorePrivate stores a newly created resource in storage.
func (pc *PostController) StorePrivate(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleStorePrivate(w, r)
	return
}

// UpdatePrivate updates the specified resource in storage.
func (pc *PostController) UpdatePrivate(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleUpdatePrivate(w, r)
	return
}

// DestroyPrivate removes the specified resource from strorage.
func (pc *PostController) DestroyPrivate(w http.ResponseWriter, r *http.Request) {
	pc.PostInteractor.HandleDestroyPrivate(w, r)
	return
}
