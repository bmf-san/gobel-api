package dto

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
)

// An AdminResponse is a success response.
type AdminResponse struct{}

// A ResponseAdmin represents the singular of admin for response.
type ResponseAdmin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A PrivateResponseAdmin represents the singular of admin for response.
type PrivateResponseAdmin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// MakeResponseHandleShowMe makes a response.
func (r *AdminResponse) MakeResponseHandleShowMe(a domain.Admin) (int, []byte, error) {
	res, err := json.Marshal(PrivateResponseAdmin{
		ID:   a.ID,
		Name: a.Name,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}
