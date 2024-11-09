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
	`

	_, err := s.db.Exec(query, tag.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s TagStore) FindAll(ctx context.Context) ([]*models.Tag, error) {
	query := `
		SELECT name
		FROM tags
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	tags := []*models.Tag{}

	for rows.Next() {
		tag := &models.Tag{}
		err := rows.Scan(&tag.Name)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
