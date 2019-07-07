package interfaces

import (
	"net/http"
)

// A JSONResponse is a presenter for jwt authentication.
type JSONResponse struct{}

// Success200 responses a success response for jwt authentication.
func (ap *JSONResponse) Success200(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

// Success201 responses a success response for jwt authentication.
func (ap *JSONResponse) Success201(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	return
}

// Error403 responses a error response 403 for JSONResponse.
func (ap *JSONResponse) Error403(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("The request is understood, but it has been refused or access is not allowed"))
	return
}

// Error404 responses a error response 403 for JSONResponse.
func (ap *JSONResponse) Error404(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("The item does ont exist"))
	return
}

// Error500 responses a error response 500 for JSONResponse.
func (ap *JSONResponse) Error500(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("An unexpected condition has occured"))
	return
}
