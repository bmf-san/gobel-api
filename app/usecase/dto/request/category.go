package request

// A IndexCategory represents the singular of IndexCategory.
type IndexCategory struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// A ShowCategoryByName represents the singular of ShowCategoryByName.
type ShowCategoryByName struct {
	Name string `json:"name"`
}

// A ShowCategoryByID represents the singular of ShowCategoryByID.
type ShowCategoryByID struct {
	ID int `json:"id"`
}

// A StoreCategory represents the singular of StoreCategory.
type StoreCategory struct {
	Name string `json:"name"`
}

// A UpdateCategory represents the singular of UpdateCategory.
type UpdateCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A DestroyCategoryByID represents the singular of DestroyCategoryByID.
type DestroyCategoryByID struct {
	ID int `json:"id"`
}
