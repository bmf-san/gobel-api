package interactor

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
	"github.com/bmf-san/goblin"
)

// A CommentInteractor is an interactor for a comment.
type CommentInteractor struct {
	CommentRepository repository.CommentRepository
	PostRepository    repository.PostRepository
	JSONResponse      dto.JSONResponse
	Logger            domain.Logger
}

// HandleIndexPrivate returns a listing of the resource.
func (ci *CommentInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := ci.CommentRepository.CountAll()
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			ci.Logger.Error(err.Error())
			ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			ci.Logger.Error(err.Error())
			ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	comments, err := ci.CommentRepository.FindAll(page, limit)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr dto.CommentResponse
	code, msg, err := cr.MakeResponseHandleIndexPrivate(comments)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	ci.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleShowPrivate display the specified resource.
func (ci *CommentInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	comment, err := ci.CommentRepository.FindByID(id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr dto.CommentResponse
	code, msg, err := cr.MakeResponseHandleShowPrivate(comment)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleStore stores a newly created resource in storage.
func (ci *CommentInteractor) HandleStore(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req dto.RequestComment
	if err = json.Unmarshal(body, &req); err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	title := goblin.GetParam(r.Context(), "title")

	post, err := ci.PostRepository.FindPublishByTitle(title)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	req.PostID = post.ID
	id, err := ci.CommentRepository.Save(req)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	comment, err := ci.CommentRepository.FindByID(id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr dto.CommentResponse
	code, msg, err := cr.MakeResponseHandleStore(comment)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleUpdateStatusPrivate updates the specified resource in storage.
func (ci *CommentInteractor) HandleUpdateStatusPrivate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req dto.RequestCommentStatus
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if err = json.Unmarshal(body, &req); err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if err = ci.CommentRepository.SaveStatusByID(req, id); err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	comment, err := ci.CommentRepository.FindByID(id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr dto.CommentResponse
	code, msg, err := cr.MakeResponseHandleUpdateStatusPrivate(comment)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, code, msg)
}
