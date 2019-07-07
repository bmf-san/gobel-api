package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/bmf-san/gobel-api/app/usecases"
)

// An JWTAuthController is a controller for an authentication.
type JWTAuthController struct {
	JWTAuthInteractor usecases.JWTAuthInteractor
}

// NewJWTAuthController creates an JWTAuthController.
func NewJWTAuthController(conn *sql.DB, logger usecases.Logger) *JWTAuthController {
	return &JWTAuthController{
		JWTAuthInteractor: usecases.JWTAuthInteractor{
			AdminRepository: &AdminRepository{
				Conn: conn,
			},
			JSONResponse: &JSONResponse{},
			Logger:       logger,
		},
	}
}

// SignIn displays a listing of the resource.
func (ac *JWTAuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	ac.JWTAuthInteractor.HandleJWTAuth(w, r)
	return
}
