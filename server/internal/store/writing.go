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
		INSERT INTO writings (type, title, content, author_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, writting.Type, writting.Title, writting.Content, writting.AuthorID).Scan(&writting.ID, &writting.CreatedAt)
}

func (s WrittingStore) FindAll(ctx context.Context) ([]*models.Writting, error) {
	query := `
		SELECT id, type, title, content, author_id, created_at, updated_at
		FROM writings
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writings := []*models.Writting{}

	for rows.Next() {
		writting := &models.Writting{}

		err := rows.Scan(&writting.ID, &writting.Type, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt)
		if err != nil {
			return nil, err
		}

		writings = append(writings, writting)
	}

	return writings, nil
}

func (s WrittingStore) FindOne(ctx context.Context, id string) (*models.Writting, error) {
	query := `
		SELECT id, type, title, content, author_id, created_at, updated_at FROM writings
	`

	row, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writting := &models.Writting{}

	if err = row.Scan(&writting.ID, &writting.Title, &writting.Type, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt); err != nil {
		return nil, err
	}

	return writting, nil
}

func (s WrittingStore) ComposeFeed(ctx context.Context, userID *string) ([]*models.Writting, error) {
	if userID == nil {

	}

	return make([]*models.Writting, 0), nil
}

func (s WrittingStore) Update(ctx context.Context, writting *models.Writting) error {
	query := `
		UPDATE writings 
		SET type = $1, title = $2, content = $3, updated_at = NOW()
		WHERE id = $4;
	`

	args := []interface{}{writting.Type, writting.Title, writting.Content, writting.ID}

	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s WrittingStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM writings 
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
