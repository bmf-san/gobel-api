package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto/request"
)

// A Tag is a repository for a post.
type Tag struct {
	ConnMySQL *sql.DB
}

// CountAll count all entities.
func (tr *Tag) CountAll() (int, error) {
	row := tr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			tags
	`)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// FindAll returns all entities.
func (tr *Tag) FindAll(page int, limit int) (domain.Tags, error) {
	var tags domain.Tags
	rows, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		ORDER BY created_at
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

	for rows.Next() {
		var (
			tagID        int
			tagName      string
			tagCreatedAt time.Time
			tagUpdatedAt time.Time
		)
		if err = rows.Scan(
			&tagID,
			&tagName,
			&tagCreatedAt,
			&tagUpdatedAt,
		); err != nil {
			return nil, err
		}
		tag := domain.Tag{
			ID:        tagID,
			Name:      tagName,
			CreatedAt: tagCreatedAt,
			UpdatedAt: tagUpdatedAt,
		}
		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

// FindByID returns the entity identified by the given id.
func (tr *Tag) FindByID(id int) (domain.Tag, error) {
	var tag domain.Tag
	row, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		WHERE
			id = ?
	`, id)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return tag, err
	}

	row.Next()
	var (
		tagID        int
		tagName      string
		tagCreatedAt time.Time
		tagUpdatedAt time.Time
	)
	if err = row.Scan(
		&tagID,
		&tagName,
		&tagCreatedAt,
		&tagUpdatedAt,
	); err != nil {
		return tag, err
	}

	return domain.Tag{
		ID:        tagID,
		Name:      tagName,
		CreatedAt: tagCreatedAt,
		UpdatedAt: tagUpdatedAt,
	}, nil
}

// FindByName returns the entity identified by the given name.
func (tr *Tag) FindByName(name string) (domain.Tag, error) {
	var tag domain.Tag
	row, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		WHERE
			name = ?
	`, name)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return tag, err
	}

	row.Next()
	var (
		tagID        int
		tagName      string
		tagCreatedAt time.Time
		tagUpdatedAt time.Time
	)
	if err = row.Scan(
		&tagID,
		&tagName,
		&tagCreatedAt,
		&tagUpdatedAt,
	); err != nil {
		return tag, err
	}

	return domain.Tag{
		ID:        tagID,
		Name:      tagName,
		CreatedAt: tagCreatedAt,
		UpdatedAt: tagUpdatedAt,
	}, nil
}

// Save saves the given entity.
func (tr *Tag) Save(req request.StoreTag) (int, error) {
	tx, err := tr.ConnMySQL.Begin()
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	now := time.Now()

	rslt, err := tx.Exec(`
		INSERT INTO
			tags(name, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`, req.Name, now, now)
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

// SaveByID save the given entity identified by the given id.
func (tr *Tag) SaveByID(req request.UpdateTag) error {
	tx, err := tr.ConnMySQL.Begin()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE tags
		SET
			name = ?,
			updated_at = ?
		WHERE id = ?
	`, req.Name, now, req.ID)
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
func (tr *Tag) DeleteByID(id int) (int, error) {
	tx, err := tr.ConnMySQL.Begin()

	row := tr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			tags
		WHERE
			id = ?
	`, id)

	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	var count int
	if err = row.Scan(&count); err != nil {
		return 0, err
	}

	_, err = tx.Exec(`
		DELETE FROM tags WHERE id = ?
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
