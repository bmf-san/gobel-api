package request

// A SignIn represents the singular of SignIn.
type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// A SignOut represents the singular of SignOut.
type SignOut struct {
	Token string `json:"token"`
}

// A Refresh represents the singular of Refresh.
type Refresh struct {
	Token string `json:"token"`
}

// A ShowUserInfo represents the singular of ShowUserInfo.
type ShowUserInfo struct {
	Token string `json:"token"`
}
