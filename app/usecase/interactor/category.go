package interactor

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A CategoryInteractor is an interactor for a category.
type CategoryInteractor struct {
	Category repository.Category
}

// Index returns a listing of the resource.
func (ci *CategoryInteractor) Index(req request.IndexCategory) (domain.Categories, Pagination, *HTTPError) {
	var cs domain.Categories
	var pn Pagination
	count, err := ci.Category.CountAll()
	if err != nil {
		return cs, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	categories, err := ci.Category.FindAll(req.Page, req.Limit)
	if err != nil {
		return cs, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return categories, pagination, nil
}

// IndexPrivate returns a listing of the resource.
func (ci *CategoryInteractor) IndexPrivate(req request.IndexCategory) (domain.Categories, Pagination, *HTTPError) {
	var cs domain.Categories
	var pn Pagination
	count, err := ci.Category.CountAll()
	if err != nil {
		return cs, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	categories, err := ci.Category.FindAll(req.Page, req.Limit)
	if err != nil {
		return cs, pn, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	pagination := pn.NewPagination(count, req.Page, req.Limit)
	return categories, pagination, nil
}

// Show display the specified resource.
func (ci *CategoryInteractor) Show(req request.ShowCategoryByName) (domain.Category, *HTTPError) {
	var c domain.Category
	category, err := ci.Category.FindByName(req.Name)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return category, nil
}

// ShowPrivate display the specified resource.
func (ci *CategoryInteractor) ShowPrivate(req request.ShowCategoryByID) (domain.Category, *HTTPError) {
	var c domain.Category
	category, err := ci.Category.FindByID(req.ID)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return category, nil
}

// StorePrivate stores a newly created resource in storage.
func (ci *CategoryInteractor) StorePrivate(req request.StoreCategory) (domain.Category, *HTTPError) {
	var c domain.Category
	id, err := ci.Category.Save(req)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	category, err := ci.Category.FindByID(id)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return category, nil
}

// UpdatePrivate updates the specified resource in storage.
func (ci *CategoryInteractor) UpdatePrivate(req request.UpdateCategory) (domain.Category, *HTTPError) {
	var c domain.Category
	if err := ci.Category.SaveByID(req); err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	category, err := ci.Category.FindByID(req.ID)
	if err != nil {
		return c, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return category, nil
}

// DestroyPrivate removes the specified resource from storage.
func (ci *CategoryInteractor) DestroyPrivate(req request.DestroyCategoryByID) *HTTPError {
	count, err := ci.Category.DeleteByID(req.ID)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
