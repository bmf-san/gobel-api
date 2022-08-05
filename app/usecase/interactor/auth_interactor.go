package interactor

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A AuthInteractor is an interactor for authentication.
type AuthInteractor struct {
	AdminRepository repository.AdminRepository
	JWTRepository   repository.JWTRepository
	JSONResponse    dto.JSONResponse
	Logger          domain.Logger
}

// HandleSignIn creates a new access token and a refresh token after password verification.
func (ai *AuthInteractor) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var req dto.RequestCredential
	err = json.Unmarshal(body, &req)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var admin domain.Admin
	admin, err = ai.AdminRepository.FindByCredentials(req)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusUnauthorized, nil)
		return
	}

	if err = admin.VerifyPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusUnauthorized, nil)
		return
	}

	var j domain.JWT
	j, err = ai.JWTRepository.SaveID(admin.ID)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusUnauthorized, nil)
		return
	}

	var ar dto.AuthResponse
	code, msg, err := ar.MakeResponseHandleSignin(j)
	if err != nil {
		ai.Logger.Error(err.Error())
	}
	ai.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleSignOut deletes an access token in storage by access uuid.
func (ai *AuthInteractor) HandleSignOut(w http.ResponseWriter, r *http.Request) {
	var j domain.JWT
	verifiedToken, err := j.GetVerifiedAccessToken(r.Header.Get("Authorization"))
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	accessUUID, err := j.GetAccessUUID(verifiedToken)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var deleted int64
	deleted, err = ai.JWTRepository.DeleteByAccessUUID(accessUUID)
	if err != nil || deleted == 0 {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	ai.JSONResponse.HTTPStatus(w, http.StatusOK, nil)
}

// HandleRefresh refreshes an access token by refresh token.
func (ai *AuthInteractor) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	var j domain.JWT
	verifiedToken, err := j.GetVerifiedRefreshToken(r.Header.Get("Authorization"))
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	refreshUUID, err := j.GetRefreshUUID(verifiedToken)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var adminID int
	adminID, err = ai.JWTRepository.FindIDByRefreshUUID(refreshUUID)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var deleted int64
	deleted, err = ai.JWTRepository.DeleteByRefreshUUID(refreshUUID)
	if err != nil || deleted == 0 {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	j, err = ai.JWTRepository.SaveID(adminID)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var ar dto.AuthResponse
	code, msg, err := ar.MakeResponseHandleRefresh(j)
	if err != nil {
		ai.Logger.Error(err.Error())
	}
	ai.JSONResponse.HTTPStatus(w, code, msg)
}

// HandleShowMe display the specified resource.
func (ai *AuthInteractor) HandleShowMe(w http.ResponseWriter, r *http.Request) {
	var j domain.JWT
	verifiedToken, err := j.GetVerifiedAccessToken(r.Header.Get("Authorization"))
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	accessUUID, err := j.GetAccessUUID(verifiedToken)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var adminID int
	adminID, err = ai.JWTRepository.FindIDByAccessUUID(accessUUID)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var admin domain.Admin
	admin, err = ai.AdminRepository.FindByID(adminID)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}

	var ar dto.AdminResponse
	code, msg, err := ar.MakeResponseHandleShowMe(admin)
	if err != nil {
		ai.Logger.Error(err.Error())
		ai.JSONResponse.HTTPStatus(w, http.StatusInternalServerError, nil)
		return
	}
	ai.JSONResponse.HTTPStatus(w, code, msg)
}
