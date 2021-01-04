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
	"github.com/go-redis/redis/v9"
)

// A PostController is a controller for a post.
type PostController struct {
	PostInteractor usecase.Post
	Logger         domain.Logger
}

// NewPostController creates a PostController.
func NewPostController(connMySQL *sql.DB, connRedis *redis.Client, logger domain.Logger) *PostController {
	return &PostController{
		PostInteractor: &interactor.PostInteractor{
			AdminRepository: &repository.AdminRepository{
				ConnMySQL: connMySQL,
				ConnRedis: connRedis,
			},
			Post: &repository.Post{
				ConnMySQL: connMySQL,
			},
			JWT: &repository.JWT{
				ConnRedis: connRedis,
			},
		},
		Logger: logger,
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexPost
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		ps, pn, herr := pc.PostInteractor.Index(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ips := response.MakeResponseIndexPost(ps)
		res, err := json.Marshal(ips)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		w.Header().Set("Pagination-Count", pn.Count)
		w.Header().Set("Pagination-Pagecount", pn.PageCount)
		w.Header().Set("Pagination-Page", pn.Page)
		w.Header().Set("Pagination-Limit", pn.Limit)
		JSONResponse(w, http.StatusOK, res)
	})
}

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexPostByName
		req.Name = goblin.GetParam(r.Context(), "name")
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		ps, pn, herr := pc.PostInteractor.IndexByCategory(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ips := response.MakeResponseIndexPost(ps)
		res, err := json.Marshal(ips)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		w.Header().Set("Pagination-Count", pn.Count)
		w.Header().Set("Pagination-Pagecount", pn.PageCount)
		w.Header().Set("Pagination-Page", pn.Page)
		w.Header().Set("Pagination-Limit", pn.Limit)
		JSONResponse(w, http.StatusOK, res)
	})
}

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexPostByName
		req.Name = goblin.GetParam(r.Context(), "name")
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		ps, pn, herr := pc.PostInteractor.IndexByTag(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ips := response.MakeResponseIndexPost(ps)
		res, err := json.Marshal(ips)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		w.Header().Set("Pagination-Count", pn.Count)
		w.Header().Set("Pagination-Pagecount", pn.PageCount)
		w.Header().Set("Pagination-Page", pn.Page)
		w.Header().Set("Pagination-Limit", pn.Limit)
		JSONResponse(w, http.StatusOK, res)
	})
}

// IndexPrivate displays a listing of the resource.
func (pc *PostController) IndexPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.IndexPost
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Limit, _ = strconv.Atoi(r.URL.Query().Get("limit"))
		ps, pn, herr := pc.PostInteractor.IndexPrivate(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		ips := response.MakeResponseIndexPost(ps)
		res, err := json.Marshal(ips)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		w.Header().Set("Pagination-Count", pn.Count)
		w.Header().Set("Pagination-Pagecount", pn.PageCount)
		w.Header().Set("Pagination-Page", pn.Page)
		w.Header().Set("Pagination-Limit", pn.Limit)
		JSONResponse(w, http.StatusOK, res)
	})
}

// Show displays the specified resource.
func (pc *PostController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowPostByTitle
		req.Title = goblin.GetParam(r.Context(), "title")
		p, herr := pc.PostInteractor.Show(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		rp := response.MakeResponseShowPost(p)
		res, err := json.Marshal(rp)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// ShowPrivate displays the specified resource.
func (pc *PostController) ShowPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req request.ShowPostByID
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = id
		p, herr := pc.PostInteractor.ShowPrivate(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		rp := response.MakeResponseShowPostPrivate(p)
		res, err := json.Marshal(rp)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// StorePrivate stores a newly created resource in storage.
func (pc *PostController) StorePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.StorePost
		err = json.Unmarshal(body, &req)
		req.Token = r.Header.Get("Authorization")
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		p, herr := pc.PostInteractor.StorePrivate(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		rp := response.MakeResponseStoreAndUpdatePostPrivate(p)
		res, err := json.Marshal(rp)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// UpdatePrivate updates the specified resource in storage.
func (pc *PostController) UpdatePrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.UpdatePost
		err = json.Unmarshal(body, &req)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.Token = r.Header.Get("Authorization")
		pid, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		req.ID = pid
		p, herr := pc.PostInteractor.UpdatePrivate(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(herr.Error()))
		}
		rp := response.MakeResponseStoreAndUpdatePostPrivate(p)
		res, err := json.Marshal(rp)
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// DestroyPrivate removes the specified resource from strorage.
func (pc *PostController) DestroyPrivate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.DestroyPostByID
		req.ID = id
		herr := pc.PostInteractor.DestroyPrivate(req)
		if herr != nil {
			pc.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.DestroyPostPrivate{
			Message: "ok",
		})
		if err != nil {
			pc.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}
