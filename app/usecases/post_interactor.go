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

// A PostInteractor is an interactor for a post.
type PostInteractor struct {
	PostRepository PostRepository
	JSONResponse   JSONResponse
	Logger         Logger
}

// HandleIndex returns a listing of the resource.
func (pi *PostInteractor) HandleIndex(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
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
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
			return
		}
	}

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublish(page, limit)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleIndex(posts))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.Success200(w, res)
	return
}

// HandleIndexByCategory returns a listing of the resource.
func (pi *PostInteractor) HandleIndexByCategory(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
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
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
			return
		}
	}

	name := goblin.GetParam(r.Context(), "name")

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublishByCategory(page, limit, name)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleIndex(posts))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.Success200(w, res)
	return
}

// HandleIndexByTag returns a listing of the resource.
func (pi *PostInteractor) HandleIndexByTag(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAllPublish()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
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
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
			return
		}
	}

	name := goblin.GetParam(r.Context(), "name")

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAllPublishByTag(page, limit, name)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleIndex(posts))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.Success200(w, res)
	return
}

// HandleIndexPrivate returns a listing of the resource.
func (pi *PostInteractor) HandleIndexPrivate(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	const defaultPage = 1
	const defaultLimit = 10

	count, err := pi.PostRepository.CountAll()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return

	}

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
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
			pi.Logger.LogError(err)
			pi.JSONResponse.Error500(w)
			return
		}
	}

	var posts domain.Posts
	posts, err = pi.PostRepository.FindAll(page, limit)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleIndexPrivate(posts))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pageCount := math.Ceil(float64(count) / float64(limit))

	w.Header().Set("Pagination-Count", strconv.Itoa(count))
	w.Header().Set("Pagination-Pagecount", fmt.Sprint(pageCount))
	w.Header().Set("Pagination-Page", strconv.Itoa(page))
	w.Header().Set("Pagination-Limit", strconv.Itoa(limit))
	pi.JSONResponse.Success200(w, res)
	return
}

// HandleShow display the specified resource.
func (pi *PostInteractor) HandleShow(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	title := goblin.GetParam(r.Context(), "title")

	var post domain.Post
	post, err := pi.PostRepository.FindByTitle(title)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleShow(post))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pi.JSONResponse.Success200(w, res)
	return
}

// HandleShowPrivate display the specified resource.
func (pi *PostInteractor) HandleShowPrivate(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var post domain.Post
	post, err = pi.PostRepository.FindByID(id)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var pr PostResponse
	var res []byte
	res, err = json.Marshal(pr.MakeResponseHandleShowPrivate(post))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pi.JSONResponse.Success200(w, res)
	return
}

// HandleStorePrivate stores a newly created resource in storage.
func (pi *PostInteractor) HandleStorePrivate(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var req RequestPost

	ja := &domain.JWTAuth{
		Token: r.Header.Get("Authorization"),
	}

	claims, err := ja.Extract()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}
	adminID, ok := claims["id"].(float64)
	if !ok {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	req.AdminID = int(adminID)
	err = json.Unmarshal(body, &req)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	err = pi.PostRepository.Save(req)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pi.JSONResponse.Success201(w, []byte("The item was created successfully"))
	return
}

// HandleUpdatePrivate updates the specified resource in storage.
func (pi *PostInteractor) HandleUpdatePrivate(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	var req RequestPost

	ja := &domain.JWTAuth{
		Token: r.Header.Get("Authorization"),
	}

	claims, err := ja.Extract()
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}
	adminID, ok := claims["id"].(float64)
	if !ok {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	req.AdminID = int(adminID)
	err = json.Unmarshal(body, &req)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}
	err = pi.PostRepository.SaveByID(req, id)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	pi.JSONResponse.Success200(w, []byte("The item was updated successfully"))
	return
}

// HandleDestroyPrivate removes the specified resource from storage.
func (pi *PostInteractor) HandleDestroyPrivate(w http.ResponseWriter, r *http.Request) {
	pi.Logger.LogAccess(r)

	id, err := strconv.Atoi(goblin.GetParam(r.Context(), "id"))
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}
	count, err := pi.PostRepository.DeleteByID(id)
	if err != nil {
		pi.Logger.LogError(err)
		pi.JSONResponse.Error500(w)
		return
	}

	if count == 0 {
		pi.JSONResponse.Error404(w)
		return
	}

	pi.JSONResponse.Success200(w, []byte("The item was deleted successfully"))
	return
}
