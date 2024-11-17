package store

import (
	"context"
	"database/sql"
	"errors"
	"rime-api/internal/models"
)

type LikeStore struct {
	db *sql.DB
}

var ErrLikeAlreadyExists = errors.New("like already exists")

func (s *LikeStore) CreateWritingLike(ctx context.Context, like *models.Like, writingID string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	exists, err := s.checkWritingLikeExists(ctx, tx, like.AuthorID, writingID)
	if err != nil {
		return err
	}
	if exists {
		return ErrLikeAlreadyExists
	}

	if err = s.createLikeTx(ctx, tx, like); err != nil {
		return err
	}
	if err = s.attachWritingLikeTx(ctx, tx, like.ID, writingID); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *LikeStore) CreateCommentLike(ctx context.Context, like *models.Like, commentID string) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	exists, err := s.checkCommentLikeExists(ctx, tx, like.AuthorID, commentID)
	if err != nil {
		return err
	}
	if exists {
		return ErrLikeAlreadyExists
	}

	// Create like and attach to comment
	if err = s.createLikeTx(ctx, tx, like); err != nil {
		return err
	}
	if err = s.attachCommentLikeTx(ctx, tx, like.ID, commentID); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *LikeStore) createLikeTx(ctx context.Context, tx *sql.Tx, like *models.Like) error {
	query := `
		INSERT INTO likes (author_id)
		VALUES ($1)
		RETURNING id
	`
	return tx.QueryRowContext(ctx, query, like.AuthorID).Scan(&like.ID)
}

func (s *LikeStore) DeleteWritingLike(ctx context.Context, authorID, writingID string) error {
	query := `
		DELETE FROM writing_likes wl
		USING likes l
		WHERE l.author_id = $1 AND wl.like_id = l.id AND wl.writing_id = $2
	`
	_, err := s.db.ExecContext(ctx, query, authorID, writingID)
	return err
}

func (s *LikeStore) DeleteCommentLike(ctx context.Context, authorID, commentID string) error {
	query := `
		
		DELETE FROM comment_likes cl
		USING likes l
		WHERE l.author_id = $1 AND cl.like_id = l.id AND cl.comment_id = $2
	`
	_, err := s.db.ExecContext(ctx, query, authorID, commentID)
	return err
}

func (s *LikeStore) attachWritingLikeTx(ctx context.Context, tx *sql.Tx, likeID, writingID string) error {
	query := `
		INSERT INTO writing_likes (like_id, writing_id)
		VALUES ($1, $2)
	`
	_, err := tx.ExecContext(ctx, query, likeID, writingID)
	return err
}

func (s *LikeStore) attachCommentLikeTx(ctx context.Context, tx *sql.Tx, likeID, commentID string) error {
	query := `
		INSERT INTO comment_likes (like_id, comment_id)
		VALUES ($1, $2)
	`
	_, err := tx.ExecContext(ctx, query, likeID, commentID)
	return err
}

func (s *LikeStore) checkWritingLikeExists(ctx context.Context, tx *sql.Tx, authorID, writingID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM writing_likes wl
			JOIN likes l ON wl.like_id = l.id
			WHERE l.author_id = $1 AND wl.writing_id = $2
		)
	`
	var exists bool
	err := tx.QueryRowContext(ctx, query, authorID, writingID).Scan(&exists)
	return exists, err
}

func (s *LikeStore) checkCommentLikeExists(ctx context.Context, tx *sql.Tx, authorID, commentID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM comment_likes cl
			JOIN likes l ON cl.like_id = l.id
			WHERE l.author_id = $1 AND cl.comment_id = $2
		)
	`
	var exists bool
	err := tx.QueryRowContext(ctx, query, authorID, commentID).Scan(&exists)
	return exists, err
}
