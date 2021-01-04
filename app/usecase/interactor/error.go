package interactor

// HTTPError is a response data for HTTP with error.
type HTTPError struct {
	Code    int
	Message string
}

// NewLogger creates a HTTPError.
func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: msg,
	}
}

func (h *HTTPError) Error() string {
	return string(h.Message)
}
