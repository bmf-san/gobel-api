package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Post is a repository for a post.
type Post struct {
	ConnMySQL *sql.DB
}

// CountAllPublic count all public entities.
func (pr *Post) CountAllPublic() (int, error) {
	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
		WHERE
			status = "public"
	`)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// CountAll count all entities.
func (pr *Post) CountAll() (int, error) {
	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
	`)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// CountAllPublicByKeyword count all public entities by keyword.
func (pr *Post) CountAllPublicByKeyword(keyword string) (int, error) {
	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
		WHERE MATCH (title, md_body)
		AGAINST (? IN BOOLEAN MODE)
		AND
			status = "public"
	`, keyword)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// CountAllPublicByCategory count all public entities by category.
func (pr *Post) CountAllPublicByCategory(name string) (int, error) {
	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
		WHERE
			category_name = ?
		AND
			status = "public"
	`, name)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// CountAllPublicByTag count all public entities by Tag.
func (pr *Post) CountAllPublicByTag(name string) (int, error) {
	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
		WHERE
			id
		IN (
			SELECT
    			post_id
			FROM
        		tags
			LEFT JOIN
        		tag_post
			ON
        		tags.id = tag_post.tag_id
			WHERE
				tags.name = ?
			)
		AND
			status = "public"
	`, name)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// FindAllPublic returns all entities.
func (pr *Post) FindAllPublic(page int, limit int) (domain.Posts, error) {
	var posts domain.Posts
	rows, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		WHERE
			status = "public"
		ORDER BY created_at
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			postID            int
			adminID           int
			categoryID        int
			postTitle         string
			postMDBody        string
			postHTMLBody      string
			postStatus        string
			postCreatedAt     time.Time
			postUpdatedAt     time.Time
			adminName         string
			adminEmail        string
			adminPassword     string
			adminCreatedAt    time.Time
			adminUpdatedAt    time.Time
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&postID,
			&adminID,
			&categoryID,
			&postTitle,
			&postMDBody,
			&postHTMLBody,
			&postStatus,
			&postCreatedAt,
			&postUpdatedAt,
			&adminName,
			&adminEmail,
			&adminPassword,
			&adminCreatedAt,
			&adminUpdatedAt,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		post := domain.Post{
			ID: postID,
			Admin: domain.Admin{
				ID:        adminID,
				Name:      adminName,
				Email:     adminEmail,
				Password:  adminPassword,
				CreatedAt: adminCreatedAt,
				UpdatedAt: adminUpdatedAt,
			},
			Category: domain.Category{
				ID:        categoryID,
				Name:      categoryName,
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
			Title:     postTitle,
			MDBody:    postMDBody,
			HTMLBody:  postHTMLBody,
			Status:    postStatus,
			CreatedAt: postCreatedAt,
			UpdatedAt: postUpdatedAt,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	postIDs := []int{}
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	queryTag := `
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id
		IN
			(%s)
	`

	var stmt string
	if len(postIDs) == 0 {
		stmt = fmt.Sprintf(queryTag, `""`)
	} else {
		stmt = fmt.Sprintf(queryTag, strings.Trim(strings.Replace(fmt.Sprint(postIDs), " ", ",", -1), "[]"))
	}

	rows, err = pr.ConnMySQL.Query(stmt)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == tagPostPostID {
				posts[p].Tags = append(posts[p].Tags, domain.Tag{
					ID:        tagID,
					Name:      tagName,
					CreatedAt: tagCreatedAt,
					UpdatedAt: tagUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
    		comments.id,
    		comments.post_id,
    		comments.body,
    		comments.status,
    		comments.created_at,
    		comments.updated_at
		FROM
    		comments
    		JOIN
        		posts
    		ON  posts.id = comments.post_id
		WHERE
    		posts.status = "public"
		AND comments.status = "public"
		ORDER BY
    		posts.id
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == commentPostID {
				posts[p].Comments = append(posts[p].Comments, domain.Comment{
					ID:        commentID,
					PostID:    commentPostID,
					Body:      commentBody,
					Status:    commentStatus,
					CreatedAt: commentCreatedAt,
					UpdatedAt: commentUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindAllPublicByKeyword returns all entities by keyword.
func (pr *Post) FindAllPublicByKeyword(page int, limit int, keyword string) (domain.Posts, error) {
	var posts domain.Posts
	rows, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		WHERE MATCH (title, md_body)
		AGAINST (? IN BOOLEAN MODE)
		AND
			status = "public"
		ORDER BY created_at
		DESC
		LIMIT ?, ?
	`, keyword, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			postID            int
			adminID           int
			categoryID        int
			postTitle         string
			postMDBody        string
			postHTMLBody      string
			postStatus        string
			postCreatedAt     time.Time
			postUpdatedAt     time.Time
			adminName         string
			adminEmail        string
			adminPassword     string
			adminCreatedAt    time.Time
			adminUpdatedAt    time.Time
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&postID,
			&adminID,
			&categoryID,
			&postTitle,
			&postMDBody,
			&postHTMLBody,
			&postStatus,
			&postCreatedAt,
			&postUpdatedAt,
			&adminName,
			&adminEmail,
			&adminPassword,
			&adminCreatedAt,
			&adminUpdatedAt,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		post := domain.Post{
			ID: postID,
			Admin: domain.Admin{
				ID:        adminID,
				Name:      adminName,
				Email:     adminEmail,
				Password:  adminPassword,
				CreatedAt: adminCreatedAt,
				UpdatedAt: adminUpdatedAt,
			},
			Category: domain.Category{
				ID:        categoryID,
				Name:      categoryName,
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
			Title:     postTitle,
			MDBody:    postMDBody,
			HTMLBody:  postHTMLBody,
			Status:    postStatus,
			CreatedAt: postCreatedAt,
			UpdatedAt: postUpdatedAt,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	postIDs := []int{}
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	queryTag := `
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id
		IN
			(%s)
	`

	var stmt string
	if len(postIDs) == 0 {
		stmt = fmt.Sprintf(queryTag, `""`)
	} else {
		stmt = fmt.Sprintf(queryTag, strings.Trim(strings.Replace(fmt.Sprint(postIDs), " ", ",", -1), "[]"))
	}

	rows, err = pr.ConnMySQL.Query(stmt)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == tagPostPostID {
				posts[p].Tags = append(posts[p].Tags, domain.Tag{
					ID:        tagID,
					Name:      tagName,
					CreatedAt: tagCreatedAt,
					UpdatedAt: tagUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
    		comments.id,
    		comments.post_id,
    		comments.body,
    		comments.status,
    		comments.created_at,
    		comments.updated_at
		FROM
    		comments
    		JOIN
        		posts
    		ON  posts.id = comments.post_id
		WHERE
    		posts.status = "public"
		AND comments.status = "public"
		ORDER BY
    		posts.id
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == commentPostID {
				posts[p].Comments = append(posts[p].Comments, domain.Comment{
					ID:        commentID,
					PostID:    commentPostID,
					Body:      commentBody,
					Status:    commentStatus,
					CreatedAt: commentCreatedAt,
					UpdatedAt: commentUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindAllPublicByCategory returns all entities by category.
func (pr *Post) FindAllPublicByCategory(page int, limit int, name string) (domain.Posts, error) {
	var posts domain.Posts
	rows, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		WHERE
			status = "public"
		AND category_name = ?
		ORDER BY created_at
		DESC
		LIMIT ?, ?
	`, name, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			postID            int
			adminID           int
			categoryID        int
			postTitle         string
			postMDBody        string
			postHTMLBody      string
			postStatus        string
			postCreatedAt     time.Time
			postUpdatedAt     time.Time
			adminName         string
			adminEmail        string
			adminPassword     string
			adminCreatedAt    time.Time
			adminUpdatedAt    time.Time
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&postID,
			&adminID,
			&categoryID,
			&postTitle,
			&postMDBody,
			&postHTMLBody,
			&postStatus,
			&postCreatedAt,
			&postUpdatedAt,
			&adminName,
			&adminEmail,
			&adminPassword,
			&adminCreatedAt,
			&adminUpdatedAt,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		post := domain.Post{
			ID: postID,
			Admin: domain.Admin{
				ID:        adminID,
				Name:      adminName,
				Email:     adminEmail,
				Password:  adminPassword,
				CreatedAt: adminCreatedAt,
				UpdatedAt: adminUpdatedAt,
			},
			Category: domain.Category{
				ID:        categoryID,
				Name:      categoryName,
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
			Title:     postTitle,
			MDBody:    postMDBody,
			HTMLBody:  postHTMLBody,
			Status:    postStatus,
			CreatedAt: postCreatedAt,
			UpdatedAt: postUpdatedAt,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	postIDs := []int{}
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	queryTag := `
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id
		IN
			(%s)
	`

	var stmt string
	if len(postIDs) == 0 {
		stmt = fmt.Sprintf(queryTag, `""`)
	} else {
		stmt = fmt.Sprintf(queryTag, strings.Trim(strings.Replace(fmt.Sprint(postIDs), " ", ",", -1), "[]"))
	}

	rows, err = pr.ConnMySQL.Query(stmt)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == tagPostPostID {
				posts[p].Tags = append(posts[p].Tags, domain.Tag{
					ID:        tagID,
					Name:      tagName,
					CreatedAt: tagCreatedAt,
					UpdatedAt: tagUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
    		comments.id,
    		comments.post_id,
    		comments.body,
    		comments.status,
    		comments.created_at,
    		comments.updated_at
		FROM
    		comments
    		JOIN
        		view_posts
    		ON  view_posts.id = comments.post_id
		WHERE
    		view_posts.status = "public"
		AND comments.status = "public"
		AND view_posts.category_name = ?
		ORDER BY
    		view_posts.id
		DESC
		LIMIT ?, ?
	`, name, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == commentPostID {
				posts[p].Comments = append(posts[p].Comments, domain.Comment{
					ID:        commentID,
					PostID:    commentPostID,
					Body:      commentBody,
					Status:    commentStatus,
					CreatedAt: commentCreatedAt,
					UpdatedAt: commentUpdatedAt,
				})
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindAllPublicByTag returns all entities by tag.
func (pr *Post) FindAllPublicByTag(page int, limit int, name string) (domain.Posts, error) {
	var posts domain.Posts
	rows, err := pr.ConnMySQL.Query(`
	SELECT
		*
	FROM
		view_posts
	WHERE
		id
	IN (
		SELECT
    		post_id
		FROM
        	tags
		LEFT JOIN
        	tag_post
		ON
        	tags.id = tag_post.tag_id
		WHERE
			tags.name = ?
		)
	ORDER BY created_at
	DESC
	LIMIT ?, ?
	`, name, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			postID            int
			adminID           int
			categoryID        int
			postTitle         string
			postMDBody        string
			postHTMLBody      string
			postStatus        string
			postCreatedAt     time.Time
			postUpdatedAt     time.Time
			adminName         string
			adminEmail        string
			adminPassword     string
			adminCreatedAt    time.Time
			adminUpdatedAt    time.Time
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&postID,
			&adminID,
			&categoryID,
			&postTitle,
			&postMDBody,
			&postHTMLBody,
			&postStatus,
			&postCreatedAt,
			&postUpdatedAt,
			&adminName,
			&adminEmail,
			&adminPassword,
			&adminCreatedAt,
			&adminUpdatedAt,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		post := domain.Post{
			ID: postID,
			Admin: domain.Admin{
				ID:        adminID,
				Name:      adminName,
				Email:     adminEmail,
				Password:  adminPassword,
				CreatedAt: adminCreatedAt,
				UpdatedAt: adminUpdatedAt,
			},
			Category: domain.Category{
				ID:        categoryID,
				Name:      categoryName,
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
			Title:     postTitle,
			MDBody:    postMDBody,
			HTMLBody:  postHTMLBody,
			Status:    postStatus,
			CreatedAt: postCreatedAt,
			UpdatedAt: postUpdatedAt,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	postIDs := []int{}
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	queryTag := `
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id
		IN
			(%s)
	`

	var stmt string
	if len(postIDs) == 0 {
		stmt = fmt.Sprintf(queryTag, `""`)
	} else {
		stmt = fmt.Sprintf(queryTag, strings.Trim(strings.Replace(fmt.Sprint(postIDs), " ", ",", -1), "[]"))
	}

	rows, err = pr.ConnMySQL.Query(stmt)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == tagPostPostID {
				posts[p].Tags = append(posts[p].Tags, domain.Tag{
					ID:        tagID,
					Name:      tagName,
					CreatedAt: tagCreatedAt,
					UpdatedAt: tagUpdatedAt,
				})
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
			comments.id,
    		comments.post_id,
    		comments.body,
    		comments.status,
    		comments.created_at,
    		comments.updated_at
		FROM
			comments
			JOIN
				posts
			ON  posts.id = comments.post_id
		WHERE
			posts.id IN(
				SELECT
					tag_post.post_id
				FROM
					tags
					LEFT JOIN
						tag_post
					ON  tags.id = tag_post.tag_id
				WHERE
					tags.name = ?
			)
		AND posts.status = "public"
		AND comments.status = "public"
		ORDER BY
			posts.id
		LIMIT ?, ?
	`, name, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == commentPostID {
				posts[p].Comments = append(posts[p].Comments, domain.Comment{
					ID:        commentID,
					PostID:    commentPostID,
					Body:      commentBody,
					Status:    commentStatus,
					CreatedAt: commentCreatedAt,
					UpdatedAt: commentUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindAll returns all entities.
func (pr *Post) FindAll(page int, limit int) (domain.Posts, error) {
	var posts domain.Posts
	rows, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		ORDER BY created_at
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			postID            int
			adminID           int
			categoryID        int
			postTitle         string
			postMDBody        string
			postHTMLBody      string
			postStatus        string
			postCreatedAt     time.Time
			postUpdatedAt     time.Time
			adminName         string
			adminEmail        string
			adminPassword     string
			adminCreatedAt    time.Time
			adminUpdatedAt    time.Time
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&postID,
			&adminID,
			&categoryID,
			&postTitle,
			&postMDBody,
			&postHTMLBody,
			&postStatus,
			&postCreatedAt,
			&postUpdatedAt,
			&adminName,
			&adminEmail,
			&adminPassword,
			&adminCreatedAt,
			&adminUpdatedAt,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		post := domain.Post{
			ID: postID,
			Admin: domain.Admin{
				ID:        adminID,
				Name:      adminName,
				Email:     adminEmail,
				Password:  adminPassword,
				CreatedAt: adminCreatedAt,
				UpdatedAt: adminUpdatedAt,
			},
			Category: domain.Category{
				ID:        categoryID,
				Name:      categoryName,
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
			Title:     postTitle,
			MDBody:    postMDBody,
			HTMLBody:  postHTMLBody,
			Status:    postStatus,
			CreatedAt: postCreatedAt,
			UpdatedAt: postUpdatedAt,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	postIDs := []int{}
	for _, p := range posts {
		postIDs = append(postIDs, p.ID)
	}

	queryTag := `
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id
		IN
			(%s)
	`

	var stmt string
	if len(postIDs) == 0 {
		stmt = fmt.Sprintf(queryTag, `""`)
	} else {
		stmt = fmt.Sprintf(queryTag, strings.Trim(strings.Replace(fmt.Sprint(postIDs), " ", ",", -1), "[]"))
	}

	rows, err = pr.ConnMySQL.Query(stmt)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == tagPostPostID {
				posts[p].Tags = append(posts[p].Tags, domain.Tag{
					ID:        tagID,
					Name:      tagName,
					CreatedAt: tagCreatedAt,
					UpdatedAt: tagUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
    		comments.id,
    		comments.post_id,
    		comments.body,
    		comments.status,
    		comments.created_at,
    		comments.updated_at
		FROM
    		comments
    		JOIN
        		posts
    		ON  posts.id = comments.post_id
		ORDER BY
    		posts.id
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return nil, err
		}

		for p := range posts {
			if posts[p].ID == commentPostID {
				posts[p].Comments = append(posts[p].Comments, domain.Comment{
					ID:        commentID,
					PostID:    commentPostID,
					Body:      commentBody,
					Status:    commentStatus,
					CreatedAt: commentCreatedAt,
					UpdatedAt: commentUpdatedAt,
				})
			}

		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// FindPublicByTitle returns the entity identified by the given title.
func (pr *Post) FindPublicByTitle(title string) (domain.Post, error) {
	var post domain.Post
	row, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		WHERE
			title = ?
		AND
			status = "public"
	`, title)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	row.Next()
	var (
		postID            int
		adminID           int
		categoryID        int
		postTitle         string
		postMDBody        string
		postHTMLBody      string
		postStatus        string
		postCreatedAt     time.Time
		postUpdatedAt     time.Time
		adminName         string
		adminEmail        string
		adminPassword     string
		adminCreatedAt    time.Time
		adminUpdatedAt    time.Time
		categoryName      string
		categoryCreatedAt time.Time
		categoryUpdatedAt time.Time
	)
	if err = row.Scan(
		&postID,
		&adminID,
		&categoryID,
		&postTitle,
		&postMDBody,
		&postHTMLBody,
		&postStatus,
		&postCreatedAt,
		&postUpdatedAt,
		&adminName,
		&adminEmail,
		&adminPassword,
		&adminCreatedAt,
		&adminUpdatedAt,
		&categoryName,
		&categoryCreatedAt,
		&categoryUpdatedAt,
	); err != nil {
		return post, nil
	}
	p := domain.Post{
		ID: postID,
		Admin: domain.Admin{
			ID:        adminID,
			Name:      adminName,
			Email:     adminEmail,
			Password:  adminPassword,
			CreatedAt: adminCreatedAt,
			UpdatedAt: adminUpdatedAt,
		},
		Category: domain.Category{
			ID:        categoryID,
			Name:      categoryName,
			CreatedAt: categoryCreatedAt,
			UpdatedAt: categoryUpdatedAt,
		},
		Title:     postTitle,
		MDBody:    postMDBody,
		HTMLBody:  postHTMLBody,
		Status:    postStatus,
		CreatedAt: postCreatedAt,
		UpdatedAt: postUpdatedAt,
	}

	rows, err := pr.ConnMySQL.Query(`
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id = ?
	`, p.ID)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return post, err
		}

		if p.ID == tagPostPostID {
			p.Tags = append(p.Tags, domain.Tag{
				ID:        tagID,
				Name:      tagName,
				CreatedAt: tagCreatedAt,
				UpdatedAt: tagUpdatedAt,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return post, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			comments
		WHERE
			status = "public"
		AND
			post_id = ?
	`, p.ID)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return post, err
		}

		if p.ID == commentPostID {
			p.Comments = append(p.Comments, domain.Comment{
				ID:        commentID,
				PostID:    commentPostID,
				Body:      commentBody,
				Status:    commentStatus,
				CreatedAt: commentCreatedAt,
				UpdatedAt: commentUpdatedAt,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return post, err
	}

	return p, nil
}

// FindByID returns the entity identified by the given id.
func (pr *Post) FindByID(id int) (domain.Post, error) {
	var post domain.Post
	row, err := pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			view_posts
		WHERE
			id = ?
	`, id)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	row.Next()
	var (
		postID            int
		adminID           int
		categoryID        int
		postTitle         string
		postMDBody        string
		postHTMLBody      string
		postStatus        string
		postCreatedAt     time.Time
		postUpdatedAt     time.Time
		adminName         string
		adminEmail        string
		adminPassword     string
		adminCreatedAt    time.Time
		adminUpdatedAt    time.Time
		categoryName      string
		categoryCreatedAt time.Time
		categoryUpdatedAt time.Time
	)
	if err = row.Scan(
		&postID,
		&adminID,
		&categoryID,
		&postTitle,
		&postMDBody,
		&postHTMLBody,
		&postStatus,
		&postCreatedAt,
		&postUpdatedAt,
		&adminName,
		&adminEmail,
		&adminPassword,
		&adminCreatedAt,
		&adminUpdatedAt,
		&categoryName,
		&categoryCreatedAt,
		&categoryUpdatedAt,
	); err != nil {
		return post, err
	}
	p := domain.Post{
		ID: postID,
		Admin: domain.Admin{
			ID:        adminID,
			Name:      adminName,
			Email:     adminEmail,
			Password:  adminPassword,
			CreatedAt: adminCreatedAt,
			UpdatedAt: adminUpdatedAt,
		},
		Category: domain.Category{
			ID:        categoryID,
			Name:      categoryName,
			CreatedAt: categoryCreatedAt,
			UpdatedAt: categoryUpdatedAt,
		},
		Title:     postTitle,
		MDBody:    postMDBody,
		HTMLBody:  postHTMLBody,
		Status:    postStatus,
		CreatedAt: postCreatedAt,
		UpdatedAt: postUpdatedAt,
	}

	rows, err := pr.ConnMySQL.Query(`
		SELECT
			tag_post.post_id AS tag_post_post_id,
			tags.id AS tag_id,
			tags.name AS tag_name,
			tags.created_at AS tag_created_at,
			tags.updated_at AS tag_updated_at
		FROM
			tags
		LEFT JOIN
			tag_post
		ON
			tags.id = tag_post.tag_id
		WHERE
			tag_post.post_id = ?
	`, p.ID)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	for rows.Next() {
		var (
			tagPostPostID int
			tagID         int
			tagName       string
			tagCreatedAt  time.Time
			tagUpdatedAt  time.Time
		)
		if err = rows.Scan(&tagPostPostID, &tagID, &tagName, &tagCreatedAt, &tagUpdatedAt); err != nil {
			return post, err
		}

		if p.ID == tagPostPostID {
			p.Tags = append(p.Tags, domain.Tag{
				ID:        tagID,
				Name:      tagName,
				CreatedAt: tagCreatedAt,
				UpdatedAt: tagUpdatedAt,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return post, err
	}

	rows, err = pr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			comments
		WHERE
			status = "public"
		AND
			post_id = ?
	`, p.ID)

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return post, err
	}

	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(&commentID, &commentPostID, &commentBody, &commentStatus, &commentCreatedAt, &commentUpdatedAt); err != nil {
			return post, err
		}

		if p.ID == commentPostID {
			p.Comments = append(p.Comments, domain.Comment{
				ID:        commentID,
				PostID:    commentPostID,
				Body:      commentBody,
				Status:    commentStatus,
				CreatedAt: commentCreatedAt,
				UpdatedAt: commentUpdatedAt,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return post, err
	}

	return p, nil
}

// Save saves the given entity.
func (pr *Post) Save(req request.StorePost) (int, error) {
	tx, err := pr.ConnMySQL.Begin()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	now := time.Now()

	res, err := tx.Exec(`
		INSERT INTO
			posts(admin_id, category_id, title, md_body, html_body, status, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?)
	`, req.AdminID, req.CategoryID, req.Title, req.MDBody, req.HTMLBody, req.Status, now, now)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	vStrings := []string{}
	vArgs := []interface{}{}
	for _, t := range req.Tags {
		vStrings = append(vStrings, "(?, ?, ?, ?)")

		vArgs = append(vArgs, t.ID)
		vArgs = append(vArgs, postID)
		vArgs = append(vArgs, now)
		vArgs = append(vArgs, now)
	}

	// bulk insert
	const queryTag = `
		INSERT INTO
			tag_post(tag_id, post_id, created_at, updated_at)
		VALUES
	  		%s
	`
	stmt := fmt.Sprintf(queryTag, strings.Join(vStrings, ","))

	_, err = tx.Exec(stmt, vArgs...)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(postID), nil
}

// SaveByID save the given entity identified by the given id.
func (pr *Post) SaveByID(req request.UpdatePost) error {
	tx, err := pr.ConnMySQL.Begin()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE posts
		SET
			admin_id = ?,
			category_id = ?,
			title = ?,
			md_body = ?,
			html_body = ?,
			status = ?,
			updated_at = ?
		WHERE id = ?
	`, req.AdminID, req.CategoryID, req.Title, req.MDBody, req.HTMLBody, req.Status, now, req.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		DELETE FROM tag_post WHERE post_id = ?
	`, req.ID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	vStrings := []string{}
	vArgs := []interface{}{}
	for _, t := range req.Tags {
		vStrings = append(vStrings, "(?, ?, ?, ?)")

		vArgs = append(vArgs, t.ID)
		vArgs = append(vArgs, req.ID)
		vArgs = append(vArgs, now)
		vArgs = append(vArgs, now)
	}

	// bulk insert
	const queryTag = `
		INSERT INTO
			tag_post(tag_id, post_id, created_at, updated_at)
		VALUES
	  		%s
	`
	stmt := fmt.Sprintf(queryTag, strings.Join(vStrings, ","))

	_, err = tx.Exec(stmt, vArgs...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes the entity identified by the given id.
func (pr *Post) DeleteByID(id int) (int, error) {
	tx, err := pr.ConnMySQL.Begin()

	row := pr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			view_posts
		WHERE
			id = ?
	`, id)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
		DELETE FROM tag_post WHERE post_id = ?
	`, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
		DELETE FROM comments WHERE post_id = ?
	`, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(`
		DELETE FROM posts WHERE id = ?
	`, id)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return count, nil
}
