package dto

// A RequestComment represents the singular of comment for store.
type RequestComment struct {
	PostID int    `json:"post_id"`
	Body   string `json:"body"`
}

// A RequestCommentStatus represents the singular of comment for store.
type RequestCommentStatus struct {
	Status string `json:"status"`
}
