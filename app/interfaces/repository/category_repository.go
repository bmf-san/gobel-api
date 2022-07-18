package repository

import (
	"database/sql"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/bmf-san/gobel-api/app/usecase/dto"
)

// A CategoryRepository is a repository for a comment.[]
type CategoryRepository struct {
	ConnMySQL *sql.DB
}

// CountAll count all entities.
func (cr *CategoryRepository) CountAll() (int, error) {
	row := cr.ConnMySQL.QueryRow(`
		SELECT
			count(*)
		FROM
			categories
	`)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// FindAll returns all entities.
func (cr *CategoryRepository) FindAll(page int, limit int) (domain.Categories, error) {
	rows, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
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

	var categories domain.Categories
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
func (cr *CategoryRepository) FindByID(id int) (domain.Category, error) {
	var category domain.Category
	row, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
		WHERE
			id = ?
	`, id)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return category, err
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
		return category, nil
	}

	return domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}, nil
}

// FindByName returns the entity identified by the given id.
func (cr *CategoryRepository) FindByName(name string) (domain.Category, error) {
	var category domain.Category
	row, err := cr.ConnMySQL.Query(`
		SELECT
			*
		FROM
			categories
		WHERE
			name = ?
	`, name)

	defer func() {
		if rerr := row.Close(); rerr != nil {
			err = rerr
		}
	}()

	if err != nil {
		return category, nil
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
		return category, nil
	}

	return domain.Category{
		ID:        categoryID,
		Name:      categoryName,
		CreatedAt: categoryCreatedAt,
		UpdatedAt: categoryUpdatedAt,
	}, nil
}

// Save saves the given entity.
func (cr *CategoryRepository) Save(req dto.RequestCategory) (int, error) {
	tx, err := cr.ConnMySQL.Begin()
	if err != nil {
		return 0, err
	}

	now := time.Now()

	rslt, err := tx.Exec(`
		INSERT INTO
			categories(name, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`, req.Name, now, now)
	if err != nil {
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
func (cr *CategoryRepository) SaveByID(req dto.RequestCategory, id int) error {
	tx, err := cr.ConnMySQL.Begin()
	if err != nil {
		return err
	}

	now := time.Now()

	_, err = tx.Exec(`
		UPDATE categories
		SET
			name = ?,
			updated_at = ?
		WHERE id = ?
	`, req.Name, now, id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return err
}

// DeleteByID deletes the entity identified by the given id.
func (cr *CategoryRepository) DeleteByID(id int) (int, error) {
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
		return 0, nil
	}

	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, nil
	}

	if count == 0 {
		return count, nil
	}

	_, err = tx.Exec(`
		DELETE FROM categories WHERE id = ?
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
