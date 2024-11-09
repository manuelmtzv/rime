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
	query := `
		INSERT INTO users (name, lastname, username, email, password) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, user.Name, user.Lastname, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
}

func (s UserStore) FindAll(ctx context.Context) ([]*models.User, error) {
	query := `
		SELECT id, name, lastname, username, email, created_at, updated_at
		FROM users
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	users := []*models.User{}

	for rows.Next() {
		user := &models.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s UserStore) FindPopular(ctx context.Context) ([]*models.PopularUser, error) {
	query := `
		SELECT users.id, users.name, users.lastname, users.username, users.email, COUNT(followers.follower_id) AS followers
		FROM users
		LEFT JOIN followers ON users.id = followers.follower_id
		GROUP BY users.id
		ORDER BY followers DESC
		LIMIT 8;		
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	popularUsers := []*models.PopularUser{}

	for rows.Next() {
		user := &models.PopularUser{}

		err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Email, &user.Followers)
		if err != nil {
			return nil, err
		}

		popularUsers = append(popularUsers, user)
	}

	return popularUsers, nil
}

func (s UserStore) FindOne(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, lastname, username, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &models.User{}

	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserStore) FindByIdentifier(ctx context.Context, identifier string) (*models.User, error) {
	query := `
		SELECT id, name, lastname, username, email, password, created_at, updated_at
		FROM users
		WHERE email = $1 OR username = $1
	`

	user := &models.User{}

	err := s.db.QueryRowContext(ctx, query, identifier).Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (s UserStore) Update(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users 
		SET name = $1, lastname = $2, username = $3, email = $4, updated_at = NOW()
		WHERE id = $5;
	`

	args := []interface{}{user.Name, user.Lastname, user.Username, user.Email, user.ID}

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
