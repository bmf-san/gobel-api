package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"log/slog"

	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/dto/response"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
	"github.com/go-redis/redis/v9"
)

// An AuthController is a controller for an authentication.
type AuthController struct {
	AuthInteractor usecase.Auth
	Logger         *slog.Logger
}

// NewAuthController creates an AuthController.
func NewAuthController(connMySQL *sql.DB, connRedis *redis.Client, logger *slog.Logger) *AuthController {
	return &AuthController{
		AuthInteractor: &interactor.AuthInteractor{
			AdminRepository: &repository.AdminRepository{
				ConnMySQL: connMySQL,
				ConnRedis: connRedis,
			},
			JWT: &repository.JWT{
				ConnRedis: connRedis,
			},
		},
		Logger: logger,
	}
}

// SignIn signs in with a credential.
func (ac *AuthController) SignIn() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		var req request.SignIn
		err = json.Unmarshal(body, &req)
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		j, herr := ac.AuthInteractor.SignIn(req)
		if herr != nil {
			ac.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.SignIn{
			AccessToken:  j.AccessToken,
			RefreshToken: j.RefreshToken,
		})
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// SignOut signs out.
func (ac *AuthController) SignOut() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		var req request.SignOut
		req.Token = t
		herr := ac.AuthInteractor.SignOut(req)
		if herr != nil {
			ac.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.SignOut{
			Message: "ok",
		})
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// Refresh refreshes a jwt.
func (ac *AuthController) Refresh() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		var req request.Refresh
		req.Token = t
		j, herr := ac.AuthInteractor.Refresh(req)
		if herr != nil {
			ac.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.Refresh{
			AccessToken:  j.AccessToken,
			RefreshToken: j.RefreshToken,
		})
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}

// ShowUserInfo displays the specified resource.
func (ac *AuthController) ShowUserInfo() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get("Authorization")
		var req request.ShowUserInfo
		req.Token = t
		a, herr := ac.AuthInteractor.ShowUserInfo(req)
		if herr != nil {
			ac.Logger.Error(herr.Error())
			JSONResponse(w, herr.Code, []byte(herr.Message))
		}
		res, err := json.Marshal(response.ShowUserInfo{
			ID:   a.ID,
			Name: a.Name,
		})
		if err != nil {
			ac.Logger.Error(err.Error())
			JSONResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		}
		JSONResponse(w, http.StatusOK, res)
	})
}
