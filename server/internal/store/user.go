package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type UserStore struct {
	db *sql.DB
}

func (s UserStore) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (s UserStore) FindOne(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, last_name, username, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &models.User{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.LastName, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserStore) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET name = $1, last_name = $2, username = $3, email = $4, updated_at = NOW()
		WHERE id = $5;
	`

	args := []interface{}{user.Name, user.LastName, user.Username, user.Email, user.ID}

	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s UserStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}