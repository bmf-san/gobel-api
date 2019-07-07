package usecases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A JWTAuthInteractor is an interactor for jwt authentication.
type JWTAuthInteractor struct {
	AdminRepository AdminRepository
	JSONResponse    JSONResponse
	Logger          Logger
}

// HandleJWTAuth returns the entity identified by authentication.
func (ai *JWTAuthInteractor) HandleJWTAuth(w http.ResponseWriter, r *http.Request) {
	ai.Logger.LogAccess(r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error500(w)
		return
	}

	var req RequestJWTAuthHandleJWTAuth
	err = json.Unmarshal(body, &req)
	if err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error500(w)
		return
	}

	var admin domain.Admin
	admin, err = ai.AdminRepository.FindByJWTAuth(req)
	if err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error500(w)
		return
	}

	var jwtAuth domain.JWTAuth
	if err = jwtAuth.SignIn(admin.Password, req.Password); err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error403(w)
		return
	}

	var token string
	token, err = jwtAuth.NewJWT(admin)
	if err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error500(w)
		return
	}

	rhja := &ResponseHandleJWTAuth{
		Token: token,
	}
	var res []byte
	res, err = json.Marshal(rhja)
	if err != nil {
		ai.Logger.LogError(err)
		ai.JSONResponse.Error500(w)
		return
	}

	ai.JSONResponse.Success200(w, res)
	return
}
