package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/goblin"
)

// A CommentInteractor is an interactor for a comment.
type CommentInteractor struct {
	CommentRepository CommentRepository
	PostRepository    PostRepository
	JSONResponse      JSONResponse
	Logger            Logger
}

// HandleIndexPrivate returns a listing of the resource.
func (ci *CommentInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	ci.Logger.LogAccess(r)

	const defaultPage = 1
	const defaultLimit = 10

	count, err := ci.CommentRepository.CountAll()
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			ci.Logger.LogError(err)
			ci.JSONResponse.Error500(w)
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
			ci.Logger.LogError(err)
			ci.JSONResponse.Error500(w)
			return
		}
	}

	var comments domain.Comments
	comments, err = ci.CommentRepository.FindAll(page, limit)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	var cr CommentResponse
	var res []byte
	res, err = json.Marshal(cr.MakeResponseHandleIndexPrivate(comments))
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	ci.JSONResponse.Success200(w, res)
	return
}

// HandleShowPrivate display the specified resource.
func (ci *CommentInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	ci.Logger.LogAccess(r)

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	var comment domain.Comment
	comment, err = ci.CommentRepository.FindByID(id)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	var cr CommentResponse
	var res []byte
	res, err = json.Marshal(cr.MakeResponseHandleShowPrivate(comment))
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	ci.JSONResponse.Success200(w, res)
	return
}

// HandleStore stores a newly created resource in storage.
func (ci *CommentInteractor) HandleStore(w http.ResponseWriter, r *http.Request) {
	ci.Logger.LogAccess(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	var req RequestComment

	err = json.Unmarshal(body, &req)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	title := goblin.GetParam(r.Context(), "title")

	var post domain.Post
	post, err = ci.PostRepository.FindByTitle(title)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	req.PostID = post.ID
	err = ci.CommentRepository.Save(req)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	ci.JSONResponse.Success201(w, []byte("The item was created successfully"))
	return
}

// HandleUpdateStatusPrivate updates the specified resource in storage.
func (ci *CommentInteractor) HandleUpdateStatusPrivate(w http.ResponseWriter, r *http.Request) {
	ci.Logger.LogAccess(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	var req RequestCommentStatus

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	err = ci.CommentRepository.SaveStatusByID(req, id)
	if err != nil {
		ci.Logger.LogError(err)
		ci.JSONResponse.Error500(w)
		return
	}

	ci.JSONResponse.Success200(w, []byte("The item was updated successfully"))
	return
}
