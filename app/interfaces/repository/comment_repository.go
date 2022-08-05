package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
)

// A CommentRepository is a repository for a comment.[]
type CommentRepository struct {
	ConnMySQL *sql.DB
}

// CountAll count all entities.
func (cr *CommentRepository) CountAll() (int, error) {
	row := cr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			comments
	`)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindAll returns all entities.
func (cr *CommentRepository) FindAll(page int, limit int) (domain.Comments, error) {
	rows, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			comments
		ORDER BY id
		DESC
		LIMIT ?, ?
	`, page*limit-limit, limit)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := rows.Close(); rerr != nil {
			err = rerr
		}
	}()

	var comments domain.Comments
	for rows.Next() {
		var (
			commentID        int
			commentPostID    int
			commentBody      string
			commentStatus    string
			commentCreatedAt time.Time
			commentUpdatedAt time.Time
		)
		if err = rows.Scan(
			&commentID,
			&commentPostID,
			&commentBody,
			&commentStatus,
			&commentCreatedAt,
			&commentUpdatedAt,
		); err != nil {
			return nil, err
		}
		comment := domain.Comment{
			ID:        commentID,
			PostID:    commentPostID,
			Body:      commentBody,
			Status:    commentStatus,
			CreatedAt: commentCreatedAt,
			UpdatedAt: commentUpdatedAt,
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// FindByID returns the entity identified by the given id.
func (cr *CommentRepository) FindByID(id int) (domain.Comment, error) {
	var comment domain.Comment
	row, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			comments
		WHERE
			id = ?
	`, id)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return comment, err
	}

	row.Next()
	var (
		commentID        int
		commentPostID    int
		commentBody      string
		commentStatus    string
		commentCreatedAt time.Time
		commentUpdatedAt time.Time
	)
	if err = row.Scan(
		&commentID,
		&commentPostID,
		&commentBody,
		&commentStatus,
		&commentCreatedAt,
		&commentUpdatedAt,
	); err != nil {
		return comment, err
	}

	return domain.Comment{
		ID:        commentID,
		PostID:    commentPostID,
		Body:      commentBody,
		Status:    commentStatus,
		CreatedAt: commentCreatedAt,
		UpdatedAt: commentUpdatedAt,
	}, nil
}

// Save saves the given entity.
func (cr *CommentRepository) Save(req dto.RequestComment) (int, error) {
	tx, err := cr.ConnMySQL.Begin()
	if err != nil {
		return 0, err
	}

	now := time.Now()

	rslt, err := tx.Exec(`
		INSERT INTO
			comments(post_id, body, created_at, updated_at)
		VALUES
			(?, ?, ?, ?)
	`, req.PostID, req.Body, now, now)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	id, err := rslt.LastInsertId()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(id), nil
}

// SaveStatusByID save the given entity identified by the given id
func (cr *CommentRepository) SaveStatusByID(req dto.RequestCommentStatus, id int) error {
	tx, err := cr.ConnMySQL.Begin()
	if err != nil {
		return err
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE comments
		SET
			status = ?,
			updated_at = ?
		WHERE id = ?
	`, req.Status, now, id)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return nil
	}

	return nil
}
