package interfaces

import (
	"database/sql"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecases"
)

// A CategoryRepository is a repository for a comment.[]
type CategoryRepository struct {
	ConnMySQL *sql.DB
}

// CountAll count all entities.
func (cr *CategoryRepository) CountAll() (count int, err error) {
	row := cr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			categories
	`)
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindAll returns all entities.
func (cr *CategoryRepository) FindAll(page int, limit int) (categories domain.Categories, err error) {
	rows, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
		ORDER BY id
		LIMIT ?, ?
	`, page*limit-limit, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			categoryID        int
			categoryName      string
			categoryCreatedAt time.Time
			categoryUpdatedAt time.Time
		)
		if err = rows.Scan(
			&categoryID,
			&categoryName,
			&categoryCreatedAt,
			&categoryUpdatedAt,
		); err != nil {
			return nil, err
		}
		category := domain.Category{
			ID:        categoryID,
			Name:      categoryName,
			CreatedAt: categoryCreatedAt,
			UpdatedAt: categoryUpdatedAt,
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// FindByID returns the entity identified by the given id.
func (cr *CategoryRepository) FindByID(id int) (category domain.Category, err error) {
	row, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
		WHERE
			id = ?
	`, id)

	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var (
		categoryID        int
		categoryName      string
		categoryCreatedAt time.Time
		categoryUpdatedAt time.Time
	)
	if err = row.Scan(
		&categoryID,
		&categoryName,
		&categoryCreatedAt,
		&categoryUpdatedAt,
	); err != nil {
		return
	}

	return domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}, nil
}

// FindByName returns the entity identified by the given id.
func (cr *CategoryRepository) FindByName(name string) (category domain.Category, err error) {
	row, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
		WHERE
			name = ?
	`, name)

	defer row.Close()

	if err != nil {
		return
	}

	row.Next()
	var (
		categoryID        int
		categoryName      string
		categoryCreatedAt time.Time
		categoryUpdatedAt time.Time
	)
	if err = row.Scan(
		&categoryID,
		&categoryName,
		&categoryCreatedAt,
		&categoryUpdatedAt,
	); err != nil {
		return
	}

	return domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}, nil
}

// Save saves the given entity.
func (cr *CategoryRepository) Save(req usecases.RequestCategory) (err error) {
	tx, err := cr.ConnMySQL.Begin()

	now := time.Now()

	_, err = tx.Exec(`
		INSERT INTO
			categories(name, created_at, updated_at)
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
func (cr *CategoryRepository) SaveByID(req usecases.RequestCategory, id int) (err error) {
	tx, err := cr.ConnMySQL.Begin()

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE categories
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
func (cr *CategoryRepository) DeleteByID(id int) (count int, err error) {
	tx, err := cr.ConnMySQL.Begin()

	row := cr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			categories
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
		DELETE FROM categories WHERE id = ?
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
