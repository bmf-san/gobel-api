package usecases

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A PostResponse is a success response.
type PostResponse struct{}

// MakeResponseHandleIndex makes a response.
func (r *PostResponse) MakeResponseHandleIndex(posts domain.Posts) ResponsePosts {
	var res []ResponsePost
	for _, p := range posts {
		responseAdmin := ResponseAdmin{
			ID:   p.Admin.ID,
			Name: p.Admin.Name,
		}
		responseCategory := ResponseCategory{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		}
		var responseTags ResponseTags
		for _, t := range p.Tags {
			responseTag := ResponseTag{
				ID:   t.ID,
				Name: t.Name,
			}
			responseTags = append(responseTags, responseTag)
		}
		var responseComments ResponseComments
		for _, c := range p.Comments {
			responseComment := ResponseComment{
				ID:        c.ID,
				Body:      c.Body,
				CreatedAt: c.CreatedAt,
			}
			responseComments = append(responseComments, responseComment)
		}
		rp := ResponsePost{
			ID:        p.ID,
			Admin:     responseAdmin,
			Category:  responseCategory,
			Tags:      responseTags,
			Title:     p.Title,
			MDBody:    p.MDBody,
			HTMLBody:  p.HTMLBody,
			Status:    p.Status,
			Comments:  responseComments,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		res = append(res, rp)
	}

	return res
}

// MakeResponseHandleIndexPrivate makes a response.
func (r *PostResponse) MakeResponseHandleIndexPrivate(posts domain.Posts) PrivateResponsePosts {
	var res []PrivateResponsePost
	for _, p := range posts {
		responseAdmin := PrivateResponseAdmin{
			ID:   p.Admin.ID,
			Name: p.Admin.Name,
		}
		responseCategory := PrivateResponseCategory{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		}
		var responseTags PrivateResponseTags
		for _, t := range p.Tags {
			responseTag := PrivateResponseTag{
				ID:   t.ID,
				Name: t.Name,
			}
			responseTags = append(responseTags, responseTag)
		}
		var responseComments PrivateResponseComments
		for _, c := range p.Comments {
			responseComment := PrivateResponseComment{
				ID:        c.ID,
				Body:      c.Body,
				CreatedAt: c.CreatedAt,
			}
			responseComments = append(responseComments, responseComment)
		}
		rp := PrivateResponsePost{
			ID:        p.ID,
			Admin:     responseAdmin,
			Category:  responseCategory,
			Tags:      responseTags,
			Title:     p.Title,
			MDBody:    p.MDBody,
			HTMLBody:  p.HTMLBody,
			Status:    p.Status,
			Comments:  responseComments,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		res = append(res, rp)
	}

	return res
}

// MakeResponseHandleShow makes a response.
func (r *PostResponse) MakeResponseHandleShow(p domain.Post) ResponsePost {
	responseAdmin := ResponseAdmin{
		ID:   p.Admin.ID,
		Name: p.Admin.Name,
	}
	responseCategory := ResponseCategory{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}
	var responseTags ResponseTags
	for _, t := range p.Tags {
		responseTag := ResponseTag{
			ID:   t.ID,
			Name: t.Name,
		}
		responseTags = append(responseTags, responseTag)
	}
	var responseComments ResponseComments
	for _, c := range p.Comments {
		responseComment := ResponseComment{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
		responseComments = append(responseComments, responseComment)
	}
	rp := ResponsePost{
		ID:        p.ID,
		Admin:     responseAdmin,
		Category:  responseCategory,
		Tags:      responseTags,
		Title:     p.Title,
		MDBody:    p.MDBody,
		HTMLBody:  p.HTMLBody,
		Status:    p.Status,
		Comments:  responseComments,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	return rp
}

// MakeResponseHandleShowPrivate makes a response.
func (r *PostResponse) MakeResponseHandleShowPrivate(p domain.Post) PrivateResponsePost {
	responseAdmin := PrivateResponseAdmin{
		ID:   p.Admin.ID,
		Name: p.Admin.Name,
	}
	responseCategory := PrivateResponseCategory{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}
	var responseTags PrivateResponseTags
	for _, t := range p.Tags {
		responseTag := PrivateResponseTag{
			ID:   t.ID,
			Name: t.Name,
		}
		responseTags = append(responseTags, responseTag)
	}
	var responseComments PrivateResponseComments
	for _, c := range p.Comments {
		responseComment := PrivateResponseComment{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
		responseComments = append(responseComments, responseComment)
	}
	rp := PrivateResponsePost{
		ID:        p.ID,
		Admin:     responseAdmin,
		Category:  responseCategory,
		Tags:      responseTags,
		Title:     p.Title,
		MDBody:    p.MDBody,
		HTMLBody:  p.HTMLBody,
		Status:    p.Status,
		Comments:  responseComments,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}

	return rp
}

// A ResponsePost represents the singular of post for response.
type ResponsePost struct {
	ID        int               `json:"id"`
	Admin     ResponseAdmin     `json:"admin"`
	Category  ResponseCategory  `json:"category"`
	Tags      []ResponseTag     `json:"tags"`
	Title     string            `json:"title"`
	MDBody    string            `json:"md_body"`
	HTMLBody  string            `json:"html_body"`
	Status    string            `json:"status"`
	Comments  []ResponseComment `json:"comments"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// A ResponsePosts represents the plural of post for response.
type ResponsePosts []ResponsePost

// A PrivateResponsePost represents the singular of post for response.
type PrivateResponsePost struct {
	ID        int                      `json:"id"`
	Admin     PrivateResponseAdmin     `json:"admin"`
	Category  PrivateResponseCategory  `json:"category"`
	Tags      []PrivateResponseTag     `json:"tags"`
	Title     string                   `json:"title"`
	MDBody    string                   `json:"md_body"`
	HTMLBody  string                   `json:"html_body"`
	Status    string                   `json:"status"`
	Comments  []PrivateResponseComment `json:"comments"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

// A PrivateResponsePosts represents the plural of post for response.
type PrivateResponsePosts []PrivateResponsePost
