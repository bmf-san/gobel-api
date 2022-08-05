package dto

// A RequestPost represents the singular of post for store.
type RequestPost struct {
	AdminID    int             `json:"admin_id"`
	CategoryID int             `json:"category_id"`
	Tags       RequestPostTags `json:"tags"`
	Title      string          `json:"title"`
	Body       string          `json:"body"`
	MDBody     string          `json:"md_body"`
	HTMLBody   string          `json:"html_body"`
	Status     string          `json:"status"`
}

// RequestPostTags represents the plural of RequestPostTag.
type RequestPostTags []RequestPostTag

// RequestPostTag represents the singular of tag for RequestPost.
type RequestPostTag struct {
	ID int `json:"id"`
}
