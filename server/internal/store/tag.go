package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"
)

type TagStore struct {
	db *sql.DB
}

func (s TagStore) Create(ctx context.Context, tag *models.Tag) error {
	query := `
		INSERT INTO tags (name) 
		VALUES ($1)
		RETURNING id, created_at
	`

	return s.db.QueryRowContext(ctx, query, tag.Name).Scan(&tag.ID, &tag.CreatedAt)
}

func (s TagStore) FindAll(ctx context.Context) ([]*models.Tag, error) {
	query := `
		SELECT id, name
		FROM tags
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	tags := []*models.Tag{}

	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (s TagStore) FindOneByName(ctx context.Context, name string) (*models.Tag, error) {
	query := `
		SELECT id, name
		FROM tags
		WHERE name = $1
	`

	tag := &models.Tag{}
	err := s.db.QueryRowContext(ctx, query, name).Scan(&tag.ID, &tag.Name)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagStore) FindPopular(ctx context.Context) ([]*models.Tag, error) {
	query := `
		SELECT id, name
		FROM tags
		LEFT JOIN tag_writing ON tags.id = tag_writing.tag_id
		GROUP BY tags.id
		ORDER BY COUNT(tag_writing.tag_id) DESC
		LIMIT 6
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	tags := []*models.Tag{}

	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
