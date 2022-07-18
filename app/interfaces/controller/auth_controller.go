package controller

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/interfaces/dto"
	"github.com/bmf-san/gobel-api/app/interfaces/repository"
	"github.com/bmf-san/gobel-api/app/usecase/interactor"
	"github.com/go-redis/redis/v9"
)

// An AuthController is a controller for an authentication.
type AuthController struct {
	AuthInteractor interactor.AuthInteractor
}

// NewAuthController creates an AuthController.
func NewAuthController(connMySQL *sql.DB, connRedis *redis.Client, logger domain.Logger) *AuthController {
	return &AuthController{
		AuthInteractor: interactor.AuthInteractor{
			AdminRepository: &repository.AdminRepository{
				ConnMySQL: connMySQL,
				ConnRedis: connRedis,
			},
			JWTRepository: &repository.JWTRepository{
				ConnRedis: connRedis,
			},
			JSONResponse: &dto.JSONResponse{},
			Logger:       logger,
		},
	}
}

// SignIn signs in with a credential.
func (ac *AuthController) SignIn() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac.AuthInteractor.HandleSignIn(w, r)
	})
}

// SignOut signs out.
func (ac *AuthController) SignOut() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac.AuthInteractor.HandleSignOut(w, r)
	})
}

// Refresh refreshes a jwt.
func (ac *AuthController) Refresh() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac.AuthInteractor.HandleRefresh(w, r)
	})
}

// ShowMe displays the specified resource.
func (ac *AuthController) ShowMe() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ac.AuthInteractor.HandleShowMe(w, r)
	})
}
