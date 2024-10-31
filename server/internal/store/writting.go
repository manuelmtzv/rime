package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type WrittingStore struct {
	db *sql.DB
}

func (s WrittingStore) Create(ctx context.Context, writting *models.Writting) error {
	query := `
		INSERT INTO writtings (type, content, author_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, writting.Type, writting.Content, writting.AuthorID).Scan(&writting.ID, &writting.CreatedAt)
}

func (s WrittingStore) FindAll(ctx context.Context) ([]*models.Writting, error) {
	query := `
		SELECT id, type, content, author_id, created_at, updated_at
		FROM writtings
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writtings := []*models.Writting{}

	for rows.Next() {
		writting := &models.Writting{}

		err := rows.Scan(&writting.ID, &writting.Type, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt)
		if err != nil {
			return nil, err
		}

		writtings = append(writtings, writting)
	}

	return writtings, nil
}

func (s WrittingStore) FindOne(ctx context.Context, id string) (*models.Writting, error) {
	query := `
		SELECT id, type, content, author_id, created_at, updated_at FROM writtings
	`

	row, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writting := &models.Writting{}

	if err = row.Scan(&writting.ID, &writting.Type, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt); err != nil {
		return nil, err
	}

	return writting, nil
}

func (s WrittingStore) Update(ctx context.Context, writting *models.Writting) error {
	query := `
		UPDATE writtings 
		SET type = $1, content = $2, updated_at = NOW()
		WHERE id = $3;
	`

	args := []interface{}{writting.Type, writting.Content, writting.ID}

	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s WrittingStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM writtings 
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
