package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecase/interactor"
)

// SetPaginationHeader sets pagination header.
func SetPaginationHeader(w http.ResponseWriter, pn interactor.Pagination) {
	w.Header().Set("Pagination-Count", pn.Count)
	w.Header().Set("Pagination-Pagecount", pn.PageCount)
	w.Header().Set("Pagination-Page", pn.Page)
	w.Header().Set("Pagination-Limit", pn.Limit)
}

// JSONResponse responses a http response with json data.
func JSONResponse(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(msg)
}
