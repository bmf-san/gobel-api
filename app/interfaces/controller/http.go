package controller

import (
	"net/http"
)

// JSONResponse responses a http response with json data.
func JSONResponse(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(msg)
}
