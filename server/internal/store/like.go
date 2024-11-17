package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type LikeStore struct {
	db *sql.DB
}

func (s LikeStore) LikeWriting(ctx context.Context, like *models.Like, writingID string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO likes (author_id)
		VALUES ($1)
		RETURNING id, created_at
	`
	if err := tx.QueryRowContext(ctx, query, like.UserID).Scan(&like.ID, &like.CreatedAt); err != nil {
		tx.Rollback()
		return err
	}

	query = `
		INSERT INTO like_writing (like_id, writing_id)
		VALUES ($1, $2)
	`
	if _, err := tx.ExecContext(ctx, query, like.ID, writingID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s LikeStore) LikeComment(ctx context.Context, like *models.Like, commentID string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO likes (author_id)
		VALUES ($1)
		RETURNING id, created_at
	`
	if err := tx.QueryRowContext(ctx, query, like.UserID).Scan(&like.ID, &like.CreatedAt); err != nil {
		tx.Rollback()
		return err
	}

	query = `
		INSERT INTO like_comment (like_id, comment_id)
		VALUES ($1, $2)
	`
	if _, err := tx.ExecContext(ctx, query, like.ID, commentID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s LikeStore) Delete(ctx context.Context, likeID string) error {
	query := `
		DELETE FROM likes
		WHERE id = $1 
	`

	_, err := s.db.ExecContext(ctx, query, likeID)
	return err
}
