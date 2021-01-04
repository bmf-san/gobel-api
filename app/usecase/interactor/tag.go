package interactor

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A TagInteractor is an interactor for a tag.
type TagInteractor struct {
	Tag repository.Tag
}

// Index returns a listing of the resource.
func (ti *TagInteractor) Index(req request.IndexTag) (domain.Tags, Pagination, *HTTPError) {
	var ts domain.Tags
	var pn Pagination
	count, err := ti.Tag.CountAll()
	if err != nil {
		return ts, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tags, err := ti.Tag.FindAll(req.Page, req.Limit)
	if err != nil {
		return ts, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return tags, pagination, nil
}

// IndexPrivate returns a listing of the resource.
func (ti *TagInteractor) IndexPrivate(req request.IndexTag) (domain.Tags, Pagination, *HTTPError) {
	var ts domain.Tags
	var pn Pagination
	count, err := ti.Tag.CountAll()
	if err != nil {
		return ts, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tags, err := ti.Tag.FindAll(req.Page, req.Limit)
	if err != nil {
		return ts, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return tags, pagination, nil
}

// Show display the specified resource.
func (ti *TagInteractor) Show(req request.ShowTagByName) (domain.Tag, *HTTPError) {
	var t domain.Tag
	tag, err := ti.Tag.FindByName(req.Name)
	if err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return tag, nil
}

// ShowPrivate display the specified resource.
func (ti *TagInteractor) ShowPrivate(req request.ShowTagByID) (domain.Tag, *HTTPError) {
	var t domain.Tag
	tag, err := ti.Tag.FindByID(req.ID)
	if err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return tag, nil
}

// StorePrivate stores a newly created resource in storage.
func (ti *TagInteractor) StorePrivate(req request.StoreTag) (domain.Tag, *HTTPError) {
	var t domain.Tag
	id, err := ti.Tag.Save(req)
	if err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tag, err := ti.Tag.FindByID(id)
	if err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return tag, nil
}

// UpdatePrivate updates the specified resource in storage.
func (ti *TagInteractor) UpdatePrivate(req request.UpdateTag) (domain.Tag, *HTTPError) {
	var t domain.Tag
	if err := ti.Tag.SaveByID(req); err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	tag, err := ti.Tag.FindByID(req.ID)
	if err != nil {
		return t, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return tag, nil
}

// DestroyPrivate removes the specified resource from storage.
func (ti *TagInteractor) DestroyPrivate(req request.DestroyTagByID) *HTTPError {
	count, err := ti.Tag.DeleteByID(req.ID)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
