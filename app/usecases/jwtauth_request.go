package usecases

// A RequestJWTAuthHandleJWTAuth represents the singular of jwtauth for jwtauth.
type RequestJWTAuthHandleJWTAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
