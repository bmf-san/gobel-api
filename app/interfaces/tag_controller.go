package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecases"
)

// A TagController is a controller for a post.
type TagController struct {
	TagInteractor usecases.TagInteractor
	Logger        usecases.Logger
}

// NewTagController creates a TagController.
func NewTagController(connMySQL *sql.DB, logger usecases.Logger) *TagController {
	return &TagController{
		TagInteractor: usecases.TagInteractor{
			TagRepository: &TagRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleIndex(w, r)
	return
}

// IndexPrivate displays a listing of the resource.
func (tc *TagController) IndexPrivate(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleIndexPrivate(w, r)
	return
}

// Show displays the specified resource.
func (tc *TagController) Show(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleShow(w, r)
	return
}

// ShowPrivate displays the specified resource.
func (tc *TagController) ShowPrivate(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleShowPrivate(w, r)
	return
}

// StorePrivate stores a newly created resource in storage.
func (tc *TagController) StorePrivate(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleStorePrivate(w, r)
	return
}

// UpdatePrivate updates the specified resource in storage.
func (tc *TagController) UpdatePrivate(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleUpdatePrivate(w, r)
	return
}

// DestroyPrivate removes the specified resource from strorage.
func (tc *TagController) DestroyPrivate(w http.ResponseWriter, r *http.Request) {
	tc.TagInteractor.HandleDestroyPrivate(w, r)
	return
}
