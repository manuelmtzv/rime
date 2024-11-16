package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type LikeStore struct {
	db *sql.DB
}

func (s LikeStore) Create(ctx context.Context, like *models.Like) error {
	query := `
		INSERT INTO likes (user_id, writing_id)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	return s.db.QueryRowContext(ctx, query, like.UserID, like.WritingID).Scan(&like.ID, &like.CreatedAt)
}

func (s LikeStore) Delete(ctx context.Context, likeID string, writingID string) error {
	query := `
		DELETE FROM likes
		WHERE id = $1 AND writing_id = $2
	`

	_, err := s.db.ExecContext(ctx, query, likeID, writingID)
	return err
}
