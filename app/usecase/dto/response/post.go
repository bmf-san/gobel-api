package response

import (
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
)

// A IndexPost represents the singular of IndexPost.
type IndexPost struct {
	ID        int           `json:"id"`
	Admin     postAdmin     `json:"admin"`
	Category  postCategory  `json:"category"`
	Tags      []postTag     `json:"tags"`
	Title     string        `json:"title"`
	MDBody    string        `json:"md_body"`
	HTMLBody  string        `json:"html_body"`
	Status    string        `json:"status"`
	Comments  []postComment `json:"comments"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// A postAdmin represents the singular of postAdmin.
type postAdmin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A postCategory represents the singular of postCategory.
type postCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A postCategoryPrivate represents the singular of postCategoryPrivate.
type postCategoryPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A postTag represents the singular of postTag.
type postTag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// A postTagPrivate represents the singular of postTagPrivate.
type postTagPrivate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A postTags represents the plural of postTag.
type postTags []postTag

// A postTagsPrivate represents the plural of postTagPrivate.
type postTagsPrivate []postTagPrivate

// A postComment represents the singular of postComment.
type postComment struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A postCommentPrivate represents the singular of postCommentPrivate.
type postCommentPrivate struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	Body      string    `json:"body"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// A postComments represents the plural of postComment.
type postComments []postComment

// A postCommentsPrivate represents the plural of postCommentPrivate.
type postCommentsPrivate []postCommentPrivate

// A ShowPost represents the singular of ShowPost.
type ShowPost struct {
	ID        int           `json:"id"`
	Admin     postAdmin     `json:"admin"`
	Category  postCategory  `json:"category"`
	Tags      []postTag     `json:"tags"`
	Title     string        `json:"title"`
	MDBody    string        `json:"md_body"`
	HTMLBody  string        `json:"html_body"`
	Status    string        `json:"status"`
	Comments  []postComment `json:"comments"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// A ShowPostPrivate represents the singular of ShowPostPrivate.
type ShowPostPrivate struct {
	ID        int                  `json:"id"`
	Admin     postAdmin            `json:"admin"`
	Category  postCategoryPrivate  `json:"category"`
	Tags      []postTagPrivate     `json:"tags"`
	Title     string               `json:"title"`
	MDBody    string               `json:"md_body"`
	HTMLBody  string               `json:"html_body"`
	Status    string               `json:"status"`
	Comments  []postCommentPrivate `json:"comments"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

// A StorePostPrivate represents the singular of StorePostPrivate.
type StorePostPrivate struct {
	ID        int                  `json:"id"`
	Admin     postAdmin            `json:"admin"`
	Category  postCategoryPrivate  `json:"category"`
	Tags      []postTagPrivate     `json:"tags"`
	Title     string               `json:"title"`
	MDBody    string               `json:"md_body"`
	HTMLBody  string               `json:"html_body"`
	Status    string               `json:"status"`
	Comments  []postCommentPrivate `json:"comments"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

// A DestroyPostPrivate represents the singular of DestroyPostPrivate.
type DestroyPostPrivate struct {
	Message string `json:"message"`
}

// MakeResponseIndexPost makes a response.
func MakeResponseIndexPost(ps domain.Posts) []IndexPost {
	var ips []IndexPost
	for _, p := range ps {
		a := postAdmin{
			ID:   p.Admin.ID,
			Name: p.Admin.Name,
		}
		c := postCategory{
			ID:   p.Category.ID,
			Name: p.Category.Name,
		}
		var ts postTags
		for _, t := range p.Tags {
			pt := postTag{
				ID:   t.ID,
				Name: t.Name,
			}
			ts = append(ts, pt)
		}
		var cms postComments
		for _, pcms := range p.Comments {
			pcm := postComment{
				ID:        pcms.ID,
				Body:      pcms.Body,
				CreatedAt: pcms.CreatedAt,
			}
			cms = append(cms, pcm)
		}
		ip := IndexPost{
			ID:        p.ID,
			Admin:     a,
			Category:  c,
			Tags:      ts,
			Title:     p.Title,
			MDBody:    p.MDBody,
			HTMLBody:  p.HTMLBody,
			Status:    p.Status,
			Comments:  cms,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		ips = append(ips, ip)
	}

	return ips

}

// MakeResponseShowPost makes a response.
func MakeResponseShowPost(p domain.Post) ShowPost {
	a := postAdmin{
		ID:   p.Admin.ID,
		Name: p.Admin.Name,
	}
	c := postCategory{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}
	var ts postTags
	for _, t := range p.Tags {
		pt := postTag{
			ID:   t.ID,
			Name: t.Name,
		}
		ts = append(ts, pt)
	}
	var cms postComments
	for _, c := range p.Comments {
		cm := postComment{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
		cms = append(cms, cm)
	}
	return ShowPost{
		ID:        p.ID,
		Admin:     a,
		Category:  c,
		Tags:      ts,
		Title:     p.Title,
		MDBody:    p.MDBody,
		HTMLBody:  p.HTMLBody,
		Status:    p.Status,
		Comments:  cms,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// MakeResponseShowPostPrivate makes a response.
func MakeResponseShowPostPrivate(p domain.Post) ShowPostPrivate {
	a := postAdmin{
		ID:   p.Admin.ID,
		Name: p.Admin.Name,
	}
	c := postCategoryPrivate{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}
	var ts postTagsPrivate
	for _, t := range p.Tags {
		pt := postTagPrivate{
			ID:   t.ID,
			Name: t.Name,
		}
		ts = append(ts, pt)
	}
	var cms postCommentsPrivate
	for _, c := range p.Comments {
		cm := postCommentPrivate{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
		cms = append(cms, cm)
	}
	return ShowPostPrivate{
		ID:        p.ID,
		Admin:     a,
		Category:  c,
		Tags:      ts,
		Title:     p.Title,
		MDBody:    p.MDBody,
		HTMLBody:  p.HTMLBody,
		Status:    p.Status,
		Comments:  cms,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// MakeResponseStoreAndUpdatePostPrivate makes a response.
func MakeResponseStoreAndUpdatePostPrivate(p domain.Post) StorePostPrivate {
	a := postAdmin{
		ID:   p.Admin.ID,
		Name: p.Admin.Name,
	}
	c := postCategoryPrivate{
		ID:   p.Category.ID,
		Name: p.Category.Name,
	}
	var ts postTagsPrivate
	for _, t := range p.Tags {
		pt := postTagPrivate{
			ID:   t.ID,
			Name: t.Name,
		}
		ts = append(ts, pt)
	}
	var cms postCommentsPrivate
	for _, c := range p.Comments {
		cm := postCommentPrivate{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
		}
		cms = append(cms, cm)
	}
	return StorePostPrivate{
		ID:        p.ID,
		Admin:     a,
		Category:  c,
		Tags:      ts,
		Title:     p.Title,
		MDBody:    p.MDBody,
		HTMLBody:  p.HTMLBody,
		Status:    p.Status,
		Comments:  cms,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
