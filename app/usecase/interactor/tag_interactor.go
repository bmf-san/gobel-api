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

// A TagInteractor is an interactor for a tag.
type TagInteractor struct {
	TagRepository repository.TagRepository
	JSONResponse  dto.JSONResponse
	Logger        domain.Logger
}

// HandleIndex returns a listing of the resource.
func (ti *TagInteractor) HandleIndex(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := ti.TagRepository.CountAll()
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			ti.Logger.Error(err.Error())
			ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
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
			ti.Logger.Error(err.Error())
			ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	tags, err := ti.TagRepository.FindAll(page, limit)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleIndex(tags)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))

	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleIndexPrivate returns a listing of the resource.
func (ti *TagInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	count, err := ti.TagRepository.CountAll()
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			ti.Logger.Error(err.Error())
			ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
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
			ti.Logger.Error(err.Error())
			ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
			return
		}
	}

	tags, err := ti.TagRepository.FindAll(page, limit)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleIndexPrivate(tags)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleShow display the specified resource.
func (ti *TagInteractor) HandleShow(w http.ResponseWriter, r *http.Request) {
	name := goblin.GetParam(r.Context(), "name")

	tag, err := ti.TagRepository.FindByName(name)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleShow(tag)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleShowPrivate display the specified resource.
func (ti *TagInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	tag, err := ti.TagRepository.FindByID(id)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleShowPrivate(tag)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleStorePrivate stores a newly created resource in storage.
func (ti *TagInteractor) HandleStorePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req dto.RequestTag
	if err = json.Unmarshal(body, &req); err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	id, err := ti.TagRepository.Save(req)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	tag, err := ti.TagRepository.FindByID(id)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleStorePrivate(tag)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleUpdatePrivate updates the specified resource in storage.
func (ti *TagInteractor) HandleUpdatePrivate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req dto.RequestTag
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if err = json.Unmarshal(body, &req); err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if err = ti.TagRepository.SaveByID(req, id); err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	tag, err := ti.TagRepository.FindByID(id)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var tr dto.TagResponse
	code, msg, err := tr.MakeResponseHandleUpdatePrivate(tag)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ti.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleDestroyPrivate removes the specified resource from storage.
func (ti *TagInteractor) HandleDestroyPrivate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	count, err := ti.TagRepository.DeleteByID(id)
	if err != nil {
		ti.Logger.Error(err.Error())
		ti.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	if count == 0 {
		ti.JSONResponse.HTTPStatus(w, http.StatusNotFound, nil)
		return
	}

	ti.JSONResponse.HTTPStatus(w, http.StatusOK, nil)
}
