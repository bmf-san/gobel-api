package interactor

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A PostInteractor is an interactor for a post.
type PostInteractor struct {
	AdminRepository repository.AdminRepository
	Post            repository.Post
	JWT             repository.JWT
}

// Index returns a listing of the resource.
func (pi *PostInteractor) Index(req request.IndexPost) (domain.Posts, Pagination, *HTTPError) {
	var ps domain.Posts
	var pn Pagination
	count, err := pi.Post.CountAllPublic()
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	posts, err := pi.Post.FindAllPublic(req.Page, req.Limit)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return posts, pagination, nil
}

// IndexByKeyword returns a listing of the resource.
func (pi *PostInteractor) IndexByKeyword(req request.IndexPostByKeyword) (domain.Posts, Pagination, *HTTPError) {
	var ps domain.Posts
	var pn Pagination
	count, err := pi.Post.CountAllPublicByKeyword(req.Keyword)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	posts, err := pi.Post.FindAllPublicByKeyword(req.Page, req.Limit, req.Keyword)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return posts, pagination, nil
}

// IndexByCategory returns a listing of the resource.
func (pi *PostInteractor) IndexByCategory(req request.IndexPostByName) (domain.Posts, Pagination, *HTTPError) {
	var ps domain.Posts
	var pn Pagination
	count, err := pi.Post.CountAllPublicByCategory(req.Name)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	posts, err := pi.Post.FindAllPublicByCategory(req.Page, req.Limit, req.Name)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return posts, pagination, nil
}

// IndexByTag returns a listing of the resource.
func (pi *PostInteractor) IndexByTag(req request.IndexPostByName) (domain.Posts, Pagination, *HTTPError) {
	var ps domain.Posts
	var pn Pagination
	count, err := pi.Post.CountAllPublicByTag(req.Name)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	posts, err := pi.Post.FindAllPublicByTag(req.Page, req.Limit, req.Name)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return posts, pagination, nil
}

// IndexPrivate returns a listing of the resource.
func (pi *PostInteractor) IndexPrivate(req request.IndexPost) (domain.Posts, Pagination, *HTTPError) {
	var ps domain.Posts
	var pn Pagination
	count, err := pi.Post.CountAll()
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	posts, err := pi.Post.FindAll(req.Page, req.Limit)
	if err != nil {
		return ps, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return posts, pagination, nil
}

// Show display the specified resource.
func (pi *PostInteractor) Show(req request.ShowPostByTitle) (domain.Post, *HTTPError) {
	var p domain.Post
	post, err := pi.Post.FindPublicByTitle(req.Title)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return post, nil
}

// ShowPrivate display the specified resource.
func (pi *PostInteractor) ShowPrivate(req request.ShowPostByID) (domain.Post, *HTTPError) {
	var p domain.Post
	post, err := pi.Post.FindByID(req.ID)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return post, nil
}

// StorePrivate stores a newly created resource in storage.
func (pi *PostInteractor) StorePrivate(req request.StorePost) (domain.Post, *HTTPError) {
	var j domain.JWT
	var p domain.Post
	vt, err := j.GetVerifiedAccessToken(req.Token)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id, err := j.GetAccessUUID(vt)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	aid, err := pi.JWT.FindIDByAccessUUID(id)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.AdminID = aid
	pid, err := pi.Post.Save(req)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	post, err := pi.Post.FindByID(pid)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return post, nil
}

// UpdatePrivate updates the specified resource in storage.
func (pi *PostInteractor) UpdatePrivate(req request.UpdatePost) (domain.Post, *HTTPError) {
	var j domain.JWT
	var p domain.Post
	vt, err := j.GetVerifiedAccessToken(req.Token)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id, err := j.GetAccessUUID(vt)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	aid, err := pi.JWT.FindIDByAccessUUID(id)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	req.AdminID = aid
	if err = pi.Post.SaveByID(req); err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	post, err := pi.Post.FindByID(req.ID)
	if err != nil {
		return p, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return post, nil
}

// DestroyPrivate removes the specified resource from storage.
func (pi *PostInteractor) DestroyPrivate(req request.DestroyPostByID) *HTTPError {
	count, err := pi.Post.DeleteByID(req.ID)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
