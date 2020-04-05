package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecases"
)

// A CategoryController is a controller for a comment.
type CategoryController struct {
	CategoryInteractor usecases.CategoryInteractor
	Logger             usecases.Logger
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(connMySQL *sql.DB, logger usecases.Logger) *CategoryController {
	return &CategoryController{
		CategoryInteractor: usecases.CategoryInteractor{
			CategoryRepository: &CategoryRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleIndex(w, r)
	return
}

// IndexPrivate displays a listing of the resource.
func (cc *CategoryController) IndexPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleIndexPrivate(w, r)
	return
}

// Show displays the specified resource.
func (cc *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleShow(w, r)
	return
}

// ShowPrivate displays the specified resource.
func (cc *CategoryController) ShowPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleShowPrivate(w, r)
	return
}

// StorePrivate stores a newly created resource in storage.
func (cc *CategoryController) StorePrivate(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleStorePrivate(w, r)
	return
}

// UpdatePrivate updates the specified resource in storage.
func (cc *CategoryController) UpdatePrivate(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleUpdatePrivate(w, r)
	return
}

// DestroyPrivate removes the specified resource from strorage.
func (cc *CategoryController) DestroyPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CategoryInteractor.HandleDestroyPrivate(w, r)
	return
}
