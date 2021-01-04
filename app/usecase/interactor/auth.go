package interactor

import (
	"net/http"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
	"github.com/bmf-san/gobel-api/app/usecase/repository"
)

// A AuthInteractor is an interactor for authentication.
type AuthInteractor struct {
	AdminRepository repository.AdminRepository
	JWT             repository.JWT
}

// SignIn creates a new access token and a refresh token after password verification.
func (ai *AuthInteractor) SignIn(req request.SignIn) (domain.JWT, *HTTPError) {
	var j domain.JWT
	var admin domain.Admin
	admin, err := ai.AdminRepository.FindByCredentials(req)
	if err != nil {
		return j, NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if err = admin.VerifyPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return j, NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	j, err = ai.JWT.SaveID(admin.ID)
	if err != nil {
		return j, NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	return j, nil
}

// SignOut deletes an access token in storage by access uuid.
func (ai *AuthInteractor) SignOut(req request.SignOut) *HTTPError {
	var j domain.JWT
	vt, err := j.GetVerifiedAccessToken(req.Token)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id, err := j.GetAccessUUID(vt)
	if err != nil {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var d int64
	d, err = ai.JWT.DeleteByAccessUUID(id)
	if err != nil || d == 0 {
		return NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

// Refresh refreshes an access token by refresh token.
func (ai *AuthInteractor) Refresh(req request.Refresh) (domain.JWT, *HTTPError) {
	var j domain.JWT
	vt, err := j.GetVerifiedRefreshToken(req.Token)
	if err != nil {
		return j, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id, err := j.GetRefreshUUID(vt)
	if err != nil {
		return j, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var aid int
	aid, err = ai.JWT.FindIDByRefreshUUID(id)
	if err != nil {
		return j, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var d int64
	d, err = ai.JWT.DeleteByRefreshUUID(id)
	if err != nil || d == 0 {
		return j, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	jr, err := ai.JWT.SaveID(aid)
	if err != nil {
		return j, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return jr, nil
}

// ShowUserInfo display the specified resource.
func (ai *AuthInteractor) ShowUserInfo(req request.ShowUserInfo) (domain.Admin, *HTTPError) {
	var j domain.JWT
	var a domain.Admin
	vt, err := j.GetVerifiedAccessToken(req.Token)
	if err != nil {
		return a, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id, err := j.GetAccessUUID(vt)
	if err != nil {
		return a, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	var aid int
	aid, err = ai.JWT.FindIDByAccessUUID(id)
	if err != nil {
		return a, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	ar, err := ai.AdminRepository.FindByID(aid)
	if err != nil {
		return a, NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ar, nil
}
