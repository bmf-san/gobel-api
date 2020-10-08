package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecase"
)

// A JSONResponse is a presenter for response.
type JSONResponse struct{}

// HTTPStatus responses a http status response for JSONResponse.
func (ap *JSONResponse) HTTPStatus(w http.ResponseWriter, code int, msg []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if msg == nil {
		msg, err := json.Marshal(usecase.ResponseHTTPStatus{
			Message: http.StatusText(code),
		})
		if err != nil {
			// NOTE: I want to log it, but consider how to do it.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
		w.Write(msg)
		return
	}

	w.Write(msg)
	return
}
