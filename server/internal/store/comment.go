package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type CommentStore struct {
	db *sql.DB
}

func (s CommentStore) Create(ctx context.Context, comment *models.Comment) error {
	query := `
		INSERT INTO comments (user_id, writing_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return s.db.QueryRowContext(ctx, query, comment.UserID, comment.WritingID, comment.Content).Scan(&comment.ID, &comment.CreatedAt)
}

func (s CommentStore) FindAll(ctx context.Context, writingID string) ([]*models.Comment, error) {
	query := `
		SELECT id, user_id, writing_id, content, created_at
		FROM comments
		WHERE writing_id = $1
		ORDER BY created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query, writingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []*models.Comment{}
	for rows.Next() {
		comment := &models.Comment{}
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.WritingID, &comment.Content, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (s CommentStore) FindOne(ctx context.Context, commentID string) (*models.Comment, error) {
	query := `
		SELECT id, user_id, writing_id, content, created_at
		FROM comments
		WHERE id = $1
	`

	comment := &models.Comment{}
	err := s.db.QueryRowContext(ctx, query, commentID).Scan(&comment.ID, &comment.UserID, &comment.WritingID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s CommentStore) Update(ctx context.Context, comment *models.Comment) error {
	query := `
		UPDATE comments
		SET content = $1
		WHERE id = $2
	`

	_, err := s.db.ExecContext(ctx, query, comment.Content, comment.ID)
	return err
}

func (s CommentStore) Delete(ctx context.Context, commentID string) error {
	query := `
		DELETE FROM comments
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, commentID)
	return err
}
