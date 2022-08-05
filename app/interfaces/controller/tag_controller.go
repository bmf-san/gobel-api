package controller

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/dto"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A TagController is a controller for a post.
type TagController struct {
	TagInteractor interactor.TagInteractor
	Logger        domain.Logger
}

// NewTagController creates a TagController.
func NewTagController(connMySQL *sql.DB, logger domain.Logger) *TagController {
	return &TagController{
		TagInteractor: interactor.TagInteractor{
			TagRepository: &repository.TagRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &dto.JSONResponse{},
			Logger:       logger,
		},
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleIndex(w, r)
	})
}

// IndexPrivate displays a listing of the resource.
func (tc *TagController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleIndexPrivate(w, r)
	})
}

// Show displays the specified resource.
func (tc *TagController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleShow(w, r)
	})
}

// ShowPrivate displays the specified resource.
func (tc *TagController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleShowPrivate(w, r)
	})
}

// StorePrivate stores a newly created resource in storage.
func (tc *TagController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleStorePrivate(w, r)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (tc *TagController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleUpdatePrivate(w, r)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (tc *TagController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.TagInteractor.HandleDestroyPrivate(w, r)
	})
}
