package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecase"
)

// A CommentController is a controller for a comment.
type CommentController struct {
	CommentInteractor usecase.CommentInteractor
	Logger            Logger
}

// NewCommentController creates a CommentController.
func NewCommentController(connMySQL *sql.DB, logger Logger) *CommentController {
	return &CommentController{
		CommentInteractor: usecase.CommentInteractor{
			CommentRepository: &CommentRepository{
				ConnMySQL: connMySQL,
			},
			PostRepository: &PostRepository{
				ConnMySQL: connMySQL,
			},
			JSONResponse: &JSONResponse{},
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
