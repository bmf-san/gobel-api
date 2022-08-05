package dto

// A RequestCredential represents the singular of credential.
type RequestCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
