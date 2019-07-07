package usecases

// A ResponseHandleJWTAuth is a success response.
type ResponseHandleJWTAuth struct {
	Token string `json:"token"`
}
