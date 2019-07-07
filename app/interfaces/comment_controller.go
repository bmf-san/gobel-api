package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecases"
)

// A CommentController is a controller for a comment.
type CommentController struct {
	CommentInteractor usecases.CommentInteractor
	Logger            usecases.Logger
}

// NewCommentController creates a CommentController.
func NewCommentController(conn *sql.DB, logger usecases.Logger) *CommentController {
	return &CommentController{
		CommentInteractor: usecases.CommentInteractor{
			CommentRepository: &CommentRepository{
				Conn: conn,
			},
			PostRepository: &PostRepository{
				Conn: conn,
			},
			JSONResponse: &JSONResponse{},
			Logger:       logger,
		},
	}
}

// IndexPrivate displays a listing of the resource.
func (cc *CommentController) IndexPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CommentInteractor.HandleIndexPrivate(w, r)
}

// ShowPrivate displays the specified resource.
func (cc *CommentController) ShowPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CommentInteractor.HandleShowPrivate(w, r)
	return
}

// Store stores a newly created resource in storage.
func (cc *CommentController) Store(w http.ResponseWriter, r *http.Request) {
	cc.CommentInteractor.HandleStore(w, r)
	return
}

// UpdateStatusPrivate updates the specified resource in storage.
func (cc *CommentController) UpdateStatusPrivate(w http.ResponseWriter, r *http.Request) {
	cc.CommentInteractor.HandleUpdateStatusPrivate(w, r)
	return
}
