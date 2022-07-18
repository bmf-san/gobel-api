package dto

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
)

// An AuthResponse is a success response.
type AuthResponse struct{}

// MakeResponseHandleSignin makes a response.
func (a *AuthResponse) MakeResponseHandleSignin(j domain.JWT) (int, []byte, error) {
	res, err := json.Marshal(ResponseSignin{
		AccessToken:  j.AccessToken,
		RefreshToken: j.RefreshToken,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// MakeResponseHandleRefresh makes a response.
func (a *AuthResponse) MakeResponseHandleRefresh(j domain.JWT) (int, []byte, error) {
	res, err := json.Marshal(ResponseRefresh{
		AccessToken:  j.AccessToken,
		RefreshToken: j.RefreshToken,
	})

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, res, nil
}

// A ResponseSignin represents the singular of signin response.
type ResponseSignin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// A ResponseRefresh represents the singular of response response.
type ResponseRefresh struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
