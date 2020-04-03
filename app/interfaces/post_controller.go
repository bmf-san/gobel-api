package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecases"
)

// A PostController is a controller for a post.
type PostController struct {
	PostInteractor usecases.PostInteractor
	Logger         usecases.Logger
}

// NewPostController creates a PostController.
func NewPostController(conn *sql.DB, logger usecases.Logger) *PostController {
	return &PostController{
		PostInteractor: usecases.PostInteractor{
			PostRepository: &PostRepository{
				Conn: conn,
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
