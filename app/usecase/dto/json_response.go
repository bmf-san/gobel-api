package dto

import "net/http"

// A JSONResponse is a presenter interface for response.
type JSONResponse interface {
	HTTPStatus(w http.ResponseWriter, code int, msg []byte)
}
