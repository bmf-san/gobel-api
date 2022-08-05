package controller

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/dto"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// A CommentController is a controller for a comment.
type CommentController struct {
	CommentInteractor interactor.CommentInteractor
	Logger            domain.Logger
}

// NewCommentController creates a CommentController.
func NewCommentController(connMySQL *sql.DB, logger domain.Logger) *CommentController {
	return &CommentController{
		CommentInteractor: interactor.CommentInteractor{
			CommentRepository: &repository.CommentRepository{
				ConnMySQL: connMySQL,
			},
			PostRepository: &repository.PostRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &dto.JSONResponse{},
			Logger:       logger,
		},
	}
}

// IndexPrivate displays a listing of the resource.
func (cc *CommentController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CommentInteractor.HandleIndexPrivate(w, r)
	})
}

// ShowPrivate displays the specified resource.
func (cc *CommentController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CommentInteractor.HandleShowPrivate(w, r)
	})
}

// Store stores a newly created resource in storage.
func (cc *CommentController) Store() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CommentInteractor.HandleStore(w, r)
	})
}

// UpdateStatusPrivate updates the specified resource in storage.
func (cc *CommentController) UpdateStatusPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.CommentInteractor.HandleUpdateStatusPrivate(w, r)
	})
}
