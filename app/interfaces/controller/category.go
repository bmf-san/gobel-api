package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/dto/response"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
	"github.com/bmf-san/goblin"
)

// A CategoryController is a controller for a comment.
type CategoryController struct {
	CategoryInteractor usecase.Category
	Logger             *slog.Logger
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(connMySQL *sql.DB, logger *slog.Logger) *CategoryController {
	return &CategoryController{
		CategoryInteractor: &interactor.CategoryInteractor{
			Category: &repository.Category{
				ConnMySQL: connMySQL,
			},
		},
		Logger: logger,
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexCategory
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		cs, pn, herr := cc.CategoryInteractor.Index(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ics := response.MakeResponseIndexCategory(cs)
		res, err := json.Marshal(ics)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		SetPaginationHeader(w, pn)
		JSONResponse(w, http.StatusOK, res)
	})
}

// IndexPrivate displays a listing of the resource.
func (cc *CategoryController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexCategory
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		cs, pn, herr := cc.CategoryInteractor.IndexPrivate(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ics := response.MakeResponseIndexCategoryPrivate(cs)
		res, err := json.Marshal(ics)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		SetPaginationHeader(w, pn)
		JSONResponse(w, http.StatusOK, res)
	})
}

// Show displays the specified resource.
func (cc *CategoryController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowCategoryByName
		req.Name = goblin.GetParam(r.Context(), "name")
		c, herr := cc.CategoryInteractor.Show(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowCategoryPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// ShowPrivate displays the specified resource.
func (cc *CategoryController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowCategoryByID
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		c, herr := cc.CategoryInteractor.ShowPrivate(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowCategoryPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// StorePrivate stores a newly created resource in storage.
func (cc *CategoryController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.StoreCategory
		err = json.Unmarshal(body, &req)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		c, herr := cc.CategoryInteractor.StorePrivate(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreCategoryPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (cc *CategoryController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.UpdateCategory
		err = json.Unmarshal(body, &req)
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		c, herr := cc.CategoryInteractor.UpdatePrivate(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		res, err := json.Marshal(response.StoreCategoryPrivate{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (cc *CategoryController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.DestroyCategoryByID
		req.ID = id
		herr := cc.CategoryInteractor.DestroyPrivate(req)
		if herr != nil {
			cc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.DestroyPostPrivate{
			Message: "ok",
		})
		if err != nil {
			cc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}
