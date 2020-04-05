package interfaces

import (
	"database/sql"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecases"
)

// A TagRepository is a repository for a post.
type TagRepository struct {
	ConnMySQL *sql.DB
}

// CountAll count all entities.
func (tr *TagRepository) CountAll() (count int, err error) {
	row := tr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			tags
	`)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindAll returns all entities.
func (tr *TagRepository) FindAll(page int, limit int) (tags domain.Tags, err error) {
	rows, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		ORDER BY id
		LIMIT ?, ?
	`, page*limit-limit, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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
func (tr *TagRepository) FindByID(id int) (tag domain.Tag, err error) {
	row, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		WHERE
			id = ?
	`, id)

	defer row.Close()

	if err != nil {
		return
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
		return
	}

	return domain.Tag{
		ID:        tagID,
		Name:      tagName,
		CreatedAt: tagCreatedAt,
		UpdatedAt: tagUpdatedAt,
	}, nil
}

// FindByName returns the entity identified by the given name.
func (tr *TagRepository) FindByName(name string) (tag domain.Tag, err error) {
	row, err := tr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			tags
		WHERE
			name = ?
	`, name)

	defer row.Close()

	if err != nil {
		return
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
		return
	}

	return domain.Tag{
		ID:        tagID,
		Name:      tagName,
		CreatedAt: tagCreatedAt,
		UpdatedAt: tagUpdatedAt,
	}, nil
}

// Save saves the given entity.
func (tr *TagRepository) Save(req usecases.RequestTag) (err error) {
	tx, err := tr.ConnMySQL.Begin()

	now := time.Now()

	_, err = tx.Exec(`
		INSERT INTO
			tags(name, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`, req.Name, now, now)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return nil
	}

	return nil
}

// SaveByID save the given entity identified by the given id.
func (tr *TagRepository) SaveByID(req usecases.RequestTag, id int) (err error) {
	tx, err := tr.ConnMySQL.Begin()

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE tags
		SET
			name = ?,
			updated_at = ?
		WHERE id = ?
	`, req.Name, now, id)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return nil
	}

	return nil
}

// DeleteByID deletes the entity identified by the given id.
func (tr *TagRepository) DeleteByID(id int) (count int, err error) {
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
		return
	}

	err = row.Scan(&count)
	if err != nil {
		return
	}

	if count == 0 {
		return count, nil
	}

	_, err = tx.Exec(`
		DELETE FROM tags WHERE id = ?
	`, id)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return count, nil
	}

	return count, nil
}
