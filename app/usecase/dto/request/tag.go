package request

// A IndexTag represents the singular of IndexTag.
type IndexTag struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// A ShowTagByName represents the singular of ShowTagByName.
type ShowTagByName struct {
	Name string `json:"name"`
}

// A ShowTagByID represents the singular of ShowTagByID.
type ShowTagByID struct {
	ID int `json:"id"`
}

// A StoreTag represents the singular of StoreTag.
type StoreTag struct {
	Name string `json:"name"`
}

// A UpdateTag represents the singular of UpdateTag.
type UpdateTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A DestroyTagByID represents the singular of DestroyTagByID.
type DestroyTagByID struct {
	ID int `json:"id"`
}
