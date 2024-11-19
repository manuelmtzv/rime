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
	query := fmt.Sprintf(`SELECT created_at FROM %s WHERE author_id = $1 AND %s = $2`, table, targetColumn)

	fmt.Println(query, authorID, targetID)

	var createdAt string
	err := app.db.QueryRowContext(ctx, query, authorID, targetID).Scan(&createdAt)

	return createdAt, err
}

func (app *LikeStore) insertLike(ctx context.Context, table string, authorID, targetID string, targetColumn string) (string, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s 
		(author_id, %s) 
		VALUES ($1, $2)
		RETURNING created_at
	`, table, targetColumn)

	var createdAt string
	err := app.db.QueryRowContext(ctx, query, authorID, targetID).Scan(&createdAt)
	return createdAt, err
}

func (app *LikeStore) deleteLike(ctx context.Context, table string, authorID, targetID string, targetColumn string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE author_id = $1 AND %s = $2`, table, targetColumn)
	_, err := app.db.ExecContext(ctx, query, authorID, targetID)
	return err
}

func (app *LikeStore) CreateWritingLike(ctx context.Context, like *models.WritingLike) error {
	createdAt, err := app.checkLikeExists(ctx, "writing_likes", like.AuthorID, like.WritingID, "writing_id")
	if err == nil && createdAt != "" {
		like.CreatedAt = createdAt
		return nil
	}

	like.CreatedAt, err = app.insertLike(ctx, "writing_likes", like.AuthorID, like.WritingID, "writing_id")
	return err
}

func (app *LikeStore) CreateCommentLike(ctx context.Context, like *models.CommentLike) error {
	createdAt, err := app.checkLikeExists(ctx, "comment_likes", like.AuthorID, like.CommentID, "comment_id")
	if err == nil && createdAt != "" {
		like.CreatedAt = createdAt
		return nil
	}
	like.CreatedAt, err = app.insertLike(ctx, "comment_likes", like.AuthorID, like.CommentID, "comment_id")
	return err
}

func (app *LikeStore) DeleteWritingLike(ctx context.Context, authorID, writingID string) error {
	return app.deleteLike(ctx, "writing_likes", authorID, writingID, "writing_id")
}

func (app *LikeStore) DeleteCommentLike(ctx context.Context, authorID, commentID string) error {
	return app.deleteLike(ctx, "comment_likes", authorID, commentID, "comment_id")
}
