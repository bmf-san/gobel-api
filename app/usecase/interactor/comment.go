package interactor

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A CommentInteractor is an interactor for a comment.
type CommentInteractor struct {
	Comment repository.Comment
	Post    repository.Post
}

// IndexPrivate returns a listing of the resource.
func (ci *CommentInteractor) IndexPrivate(req request.IndexComment) (domain.Comments, Pagination, *HTTPError) {
	var c domain.Comments
	var pn Pagination
	count, err := ci.Comment.CountAll()
	if err != nil {
		return c, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	comments, err := ci.Comment.FindAll(req.Page, req.Limit)
	if err != nil {
		return c, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return comments, pagination, nil
}

// ShowPrivate display the specified resource.
func (ci *CommentInteractor) ShowPrivate(req request.ShowCommentByID) (domain.Comment, *HTTPError) {
	var c domain.Comment
	comment, err := ci.Comment.FindByID(req.ID)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return comment, nil
}

// Store stores a newly created resource in storage.
func (ci *CommentInteractor) Store(req request.StoreComment) (domain.Comment, *HTTPError) {
	var c domain.Comment
	id, err := ci.Comment.Save(req)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	comment, err := ci.Comment.FindByID(id)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return comment, nil
}

// UpdateStatusPrivate updates the specified resource in storage.
func (ci *CommentInteractor) UpdateStatusPrivate(req request.UpdateCommentStatus) (domain.Comment, *HTTPError) {
	var c domain.Comment
	if err := ci.Comment.SaveStatusByID(req); err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	comment, err := ci.Comment.FindByID(req.ID)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return comment, nil
}
