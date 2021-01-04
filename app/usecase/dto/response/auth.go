package response

// A SignIn represents the singular of SignIn.
type SignIn struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// A SignOut represents the singular of SignOut.
type SignOut struct {
	Message string `json:"message"`
}

// A Refresh represents the singular of Refresh.
type Refresh struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// A ShowUserInfo represents the singular of ShowUserinfo.
type ShowUserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
