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

// A TagController is a controller for a post.
type TagController struct {
	TagInteractor usecase.Tag
	Logger        domain.Logger
}

// NewTagController creates a TagController.
func NewTagController(connMySQL *sql.DB, logger domain.Logger) *TagController {
	return &TagController{
		TagInteractor: &interactor.TagInteractor{
			Tag: &repository.Tag{
				ConnMySQL: connMySQL,
			},
		},
		Logger: logger,
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexTag
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		cs, pn, herr := tc.TagInteractor.Index(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ics := response.MakeResponseIndexTag(cs)
		res, err := json.Marshal(ics)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		SetPaginationHeader(w, pn)
		JSONResponse(w, http.StatusOK, res)
	})
}

// IndexPrivate displays a listing of the resource.
func (tc *TagController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexTag
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		cs, pn, herr := tc.TagInteractor.IndexPrivate(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ics := response.MakeResponseIndexTagPrivate(cs)
		res, err := json.Marshal(ics)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		SetPaginationHeader(w, pn)
		JSONResponse(w, http.StatusOK, res)
	})
}

// Show displays the specified resource.
func (tc *TagController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowTagByName
		req.Name = goblin.GetParam(r.Context(), "name")
		c, herr := tc.TagInteractor.Show(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowTagPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// ShowPrivate displays the specified resource.
func (tc *TagController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowTagByID
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		c, herr := tc.TagInteractor.ShowPrivate(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowTagPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// StorePrivate stores a newly created resource in storage.
func (tc *TagController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.StoreTag
		err = json.Unmarshal(body, &req)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		c, herr := tc.TagInteractor.StorePrivate(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreTagPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (tc *TagController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.UpdateTag
		err = json.Unmarshal(body, &req)
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		c, herr := tc.TagInteractor.UpdatePrivate(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreTagPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (tc *TagController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.DestroyTagByID
		req.ID = id
		herr := tc.TagInteractor.DestroyPrivate(req)
		if herr != nil {
			tc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.DestroyPostPrivate{
			Message: "ok",
		})
		if err != nil {
			tc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}
