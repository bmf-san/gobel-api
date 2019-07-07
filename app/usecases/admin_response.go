package usecases

// A ResponseAdmin represents the singular of admin for response.
type ResponseAdmin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A PrivateResponseAdmin represents the singular of admin for response.
type PrivateResponseAdmin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
