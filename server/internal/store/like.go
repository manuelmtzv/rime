package store

import (
	"context"
	"database/sql"
	"fmt"
	"rime-api/internal/models"
)

type LikeStore struct {
	db *sql.DB
}

func (app *LikeStore) checkLikeExists(ctx context.Context, table string, authorID, targetID string, targetColumn string) (string, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE author_id = $1 AND %s = $2`, table, targetColumn)
	var id string
	err := app.db.QueryRowContext(ctx, query, authorID, targetID).Scan(&id)
	return id, err
}

func (app *LikeStore) insertLike(ctx context.Context, table string, likeID, authorID, targetID string, createdAt string, targetColumn string) (string, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(id, author_id, %s, created_at) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`, table, targetColumn)

	var id string
	err := app.db.QueryRowContext(ctx, query, likeID, authorID, targetID, createdAt).Scan(&id)
	return id, err
}

func (app *LikeStore) deleteLike(ctx context.Context, table string, authorID, targetID string, targetColumn string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE author_id = $1 AND %s = $2`, table, targetColumn)
	_, err := app.db.ExecContext(ctx, query, authorID, targetID)
	return err
}

func (app *LikeStore) CreateWritingLike(ctx context.Context, like *models.WritingLike) error {
	id, err := app.checkLikeExists(ctx, "writing_likes", like.AuthorID, like.WritingID, "writing_id")
	if err == nil && id != "" {
		like.ID = id
		return nil
	}
	like.ID, err = app.insertLike(ctx, "writing_likes", like.ID, like.AuthorID, like.WritingID, like.CreatedAt, "writing_id")
	return err
}

func (app *LikeStore) CreateCommentLike(ctx context.Context, like *models.CommentLike) error {
	id, err := app.checkLikeExists(ctx, "comment_likes", like.AuthorID, like.CommentID, "comment_id")
	if err == nil && id != "" {
		like.ID = id
		return nil
	}
	like.ID, err = app.insertLike(ctx, "comment_likes", like.ID, like.AuthorID, like.CommentID, like.CreatedAt, "comment_id")
	return err
}

func (app *LikeStore) DeleteWritingLike(ctx context.Context, authorID, writingID string) error {
	return app.deleteLike(ctx, "writing_likes", authorID, writingID, "writing_id")
}

func (app *LikeStore) DeleteCommentLike(ctx context.Context, authorID, commentID string) error {
	return app.deleteLike(ctx, "comment_likes", authorID, commentID, "comment_id")
}
