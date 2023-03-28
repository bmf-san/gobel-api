package request

// A IndexPost represents the singular of IndexPost.
type IndexPost struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// A IndexPostByName represents the singular of IndexPostByName.
type IndexPostByName struct {
	Name  string `json:"name"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}

// A ShowPostByTitle represents the singular of ShowPostByTitle.
type ShowPostByTitle struct {
	Title string `json:"title"`
}

// A ShowPostByID represents the singular of ShowPostByID.
type ShowPostByID struct {
	ID int `json:"id"`
}

// A StorePost represents the singular of StorePost.
type StorePost struct {
	Token      string   `json:"token"`
	AdminID    int      `json:"admin_id"`
	CategoryID int      `json:"category_id"`
	Tags       postTags `json:"tags"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	MDBody     string   `json:"md_body"`
	HTMLBody   string   `json:"html_body"`
	Status     string   `json:"status"`
}

// A UpdatePost represents the singular of UpdatePost.
type UpdatePost struct {
	ID         int      `json:"id"`
	Token      string   `json:"token"`
	AdminID    int      `json:"admin_id"`
	CategoryID int      `json:"category_id"`
	Tags       postTags `json:"tags"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	MDBody     string   `json:"md_body"`
	HTMLBody   string   `json:"html_body"`
	Status     string   `json:"status"`
}

// A postTags represents the plural of postTag.
type postTags []postTag

// A postTag represents the singular of tag for postTag.
type postTag struct {
	ID int `json:"id"`
}

// A DestroyPostByID represents the singular of DestroyPostByID.
type DestroyPostByID struct {
	ID int `json:"id"`
}
