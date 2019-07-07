package usecases

import "net/http"

// A Logger represents a logger.
type Logger interface {
	LogError(e error)
	LogAccess(r *http.Request)
}
