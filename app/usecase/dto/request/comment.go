package request

// A IndexComment represents the singular of IndexComment.
type IndexComment struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// A ShowCommentByID represents the singular of ShowCommentByID.
type ShowCommentByID struct {
	ID int `json:"id"`
}

// A StoreComment represents the singular of comment for StoreComment.
type StoreComment struct {
	PostID int    `json:"post_id"`
	Body   string `json:"body"`
}

// A UpdateCommentStatus represents the singular of UpdateCommentStatus.
type UpdateCommentStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}
