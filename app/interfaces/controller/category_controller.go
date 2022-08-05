package controller

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/dto"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A CategoryController is a controller for a comment.
type CategoryController struct {
	CategoryInteractor interactor.CategoryInteractor
	Logger             domain.Logger
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(connMySQL *sql.DB, logger domain.Logger) *CategoryController {
	return &CategoryController{
		CategoryInteractor: interactor.CategoryInteractor{
			CategoryRepository: &repository.CategoryRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &dto.JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleIndex(w, r)
	})
}

// IndexPrivate displays a listing of the resource.
func (cc *CategoryController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleIndexPrivate(w, r)
	})
}

// Show displays the specified resource.
func (cc *CategoryController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleShow(w, r)
	})
}

// ShowPrivate displays the specified resource.
func (cc *CategoryController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleShowPrivate(w, r)
	})
}

// StorePrivate stores a newly created resource in storage.
func (cc *CategoryController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleStorePrivate(w, r)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (cc *CategoryController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleUpdatePrivate(w, r)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (cc *CategoryController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CategoryInteractor.HandleDestroyPrivate(w, r)
	})
}
