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

// A CategoryInteractor is an interactor for a category.
type CategoryInteractor struct {
	CategoryRepository CategoryRepository
	JSONResponse       JSONResponse
	Logger             Logger
}

// HandleIndex returns a listing of the resource.
func (ci *CategoryInteractor) HandleIndex(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := ci.CategoryRepository.CountAll()
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

	var categories domain.Categories
	categories, err = ci.CategoryRepository.FindAll(page, limit)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr CategoryResponse
	code, msg, err := cr.MakeResponseHandleIndex(categories)
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
	return
}

// HandleIndexPrivate returns a listing of the resource.
func (ci *CategoryInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := ci.CategoryRepository.CountAll()
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

	var categories domain.Categories
	categories, err = ci.CategoryRepository.FindAll(page, limit)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr CategoryResponse
	code, msg, err := cr.MakeResponseHandleIndexPrivate(categories)
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
	return
}

// HandleShow display the specified resource.
func (ci *CategoryInteractor) HandleShow(w http.ResponseWriter, r *http.Request) {
	name := goblin.GetParam(r.Context(), "name")

	var category domain.Category
	category, err := ci.CategoryRepository.FindByName(name)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr CategoryResponse
	code, msg, err := cr.MakeResponseHandleShow(category)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleShowPrivate display the specified resource.
func (ci *CategoryInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var category domain.Category
	category, err = ci.CategoryRepository.FindByID(id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var cr CategoryResponse
	code, msg, err := cr.MakeResponseHandleShowPrivate(category)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, code, msg)
	return
}

// HandleStorePrivate stores a newly created resource in storage.
func (ci *CategoryInteractor) HandleStorePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req RequestCategory

	err = json.Unmarshal(body, &req)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	err = ci.CategoryRepository.Save(req)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, http.StatusCreated, nil)
	return
}

// HandleUpdatePrivate updates the specified resource in storage.
func (ci *CategoryInteractor) HandleUpdatePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req RequestCategory

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	err = ci.CategoryRepository.SaveByID(req, id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, http.StatusOK, nil)
	return
}

// HandleDestroyPrivate removes the specified resource from storage.
func (ci *CategoryInteractor) HandleDestroyPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	count, err := ci.CategoryRepository.DeleteByID(id)
	if err != nil {
		ci.Logger.Error(err.Error())
		ci.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		ci.JSONResponse.HTTPStatus(w, http.StatusNotFound, nil)
		return
	}

	ci.JSONResponse.HTTPStatus(w, http.StatusOK, nil)
	return
}
