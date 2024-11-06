package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type WritingStore struct {
	db *sql.DB
}

func (s WritingStore) Create(ctx context.Context, writting *models.Writing) error {
	query := `
		INSERT INTO writings (type, title, content, author_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, writting.Type, writting.Title, writting.Content, writting.AuthorID).Scan(&writting.ID, &writting.CreatedAt)
}

func (s WritingStore) FindAll(ctx context.Context) ([]*models.Writing, error) {
	query := `
		SELECT id, type, title, content, author_id, created_at, updated_at
		FROM writings
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writings := []*models.Writing{}

	for rows.Next() {
		writting := &models.Writing{}

		err := rows.Scan(&writting.ID, &writting.Type, &writting.Title, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt)
		if err != nil {
			return nil, err
		}

		writings = append(writings, writting)
	}

	return writings, nil
}

func (s WritingStore) FindOne(ctx context.Context, id string) (*models.Writing, error) {
	query := `
		SELECT id, type, title, content, author_id, created_at, updated_at FROM writings
	`

	row, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	writting := &models.Writing{}

	if err = row.Scan(&writting.ID, &writting.Title, &writting.Type, &writting.Content, &writting.AuthorID, &writting.CreatedAt, &writting.UpdatedAt); err != nil {
		return nil, err
	}

	return writting, nil
}

func (s WritingStore) ComposeFeed(ctx context.Context, userID *string) ([]*models.Writing, error) {
	if userID == nil {

	}

	return make([]*models.Writing, 0), nil
}

func (s WritingStore) Update(ctx context.Context, writting *models.Writing) error {
	query := `
		UPDATE writings 
		SET type = $1, title = $2, content = $3, updated_at = NOW()
		WHERE id = $4;
	`

	args := []interface{}{writting.Type, writting.Title, writting.Content, writting.ID}

	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s WritingStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM writings 
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}
