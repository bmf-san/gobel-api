package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/dto/response"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
	"github.com/bmf-san/goblin"
)

// A CommentController is a controller for a comment.
type CommentController struct {
	CommentInteractor usecase.Comment
	Logger            domain.Logger
}

// NewCommentController creates a CommentController.
func NewCommentController(connMySQL *sql.DB, logger domain.Logger) *CommentController {
	return &CommentController{
		CommentInteractor: &interactor.CommentInteractor{
			Comment: &repository.Comment{
				ConnMySQL: connMySQL,
			},
			Post: &repository.Post{
				ConnMySQL: connMySQL,
			},
		},
		Logger: logger,
	}
}

// IndexPrivate displays a listing of the resource.
func (cc *CommentController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexComment
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		c, pn, herr := cc.CommentInteractor.IndexPrivate(req)
		if herr != nil {
			cc.Logger.ErrorContext(r.Context(), herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ics := response.MakeResponseIndexCommentPrivate(c)
		res, err := json.Marshal(ics)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		SetPaginationHeader(w, pn)
		JSONResponse(w, http.StatusOK, res)
	})
}

// ShowPrivate displays the specified resource.
func (cc *CommentController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowCommentByID
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		c, herr := cc.CommentInteractor.ShowPrivate(req)
		if herr != nil {
			cc.Logger.ErrorContext(r.Context(), herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowCommentPrivate{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// Store stores a newly created resource in storage.
func (cc *CommentController) Store() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.StoreComment
		err = json.Unmarshal(body, &req)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		c, herr := cc.CommentInteractor.Store(req)
		if herr != nil {
			cc.Logger.ErrorContext(r.Context(), herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreAndUpdateComment{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// UpdateStatusPrivate updates the specified resource in storage.
func (cc *CommentController) UpdateStatusPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.UpdateCommentStatus
		err = json.Unmarshal(body, &req)
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		c, herr := cc.CommentInteractor.UpdateStatusPrivate(req)
		if herr != nil {
			cc.Logger.ErrorContext(r.Context(), herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreAndUpdateComment{
			ID:        c.ID,
			PostID:    c.PostID,
			Body:      c.Body,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.ErrorContext(r.Context(), err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}
