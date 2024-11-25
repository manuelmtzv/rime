package store

import (
	"context"
	"database/sql"
	"rime-api/internal/models"

	"github.com/lib/pq"
)

type WritingStore struct {
	db *sql.DB
}

func (s WritingStore) Create(ctx context.Context, writting *models.Writing) error {
	query := `
		INSERT INTO writings (type, title, content, author_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id, created_at`

	return s.db.QueryRowContext(ctx, query, writting.Type, writting.Title, writting.Content, writting.AuthorID).Scan(&writting.ID, &writting.CreatedAt)
}

func (s WritingStore) FindAll(ctx context.Context) ([]*models.Writing, error) {
	query := `
		SELECT 
			writings.id, writings.type, writings.title, writings.content, 
			writings.author_id, writings.created_at, writings.updated_at,
			users.id, users.name, users.lastname, users.email, count(writing_likes.author_id) as like_count
		FROM writings
		LEFT JOIN users ON writings.author_id = users.id
		LEFT JOIN writing_likes ON writings.id = writing_likes.writing_id
		GROUP BY writings.id, users.id
		ORDER BY writings.created_at DESC
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	writings := []*models.Writing{}
	writingsIDs := []string{}

	for rows.Next() {
		writing := &models.Writing{}
		author := &models.User{}

		err := rows.Scan(
			&writing.ID,
			&writing.Type,
			&writing.Title,
			&writing.Content,
			&writing.AuthorID,
			&writing.CreatedAt,
			&writing.UpdatedAt,
			&author.ID,
			&author.Name,
			&author.Lastname,
			&author.Email,
			&writing.LikeCount,
		)
		if err != nil {
			return nil, err
		}

		writing.Author = author
		writings = append(writings, writing)
		writingsIDs = append(writingsIDs, writing.ID)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	likesMap, err := s.fetchLikesForWritings(ctx, writingsIDs)
	if err != nil {
		return nil, err
	}

	for _, writing := range writings {
		writing.Likes = likesMap[writing.ID]
	}

	return writings, nil
}

func (s WritingStore) FindOne(ctx context.Context, id string) (*models.Writing, error) {
	query := `
		SELECT writings.id, writings.type, writings.title, writings.content, writings.author_id, writings.created_at, writings.updated_at, users.id AS author_id, users.name, users.lastname, users.email
		FROM writings 
		LEFT JOIN users ON writings.author_id = users.id
		WHERE writings.id = $1
	`

	writing := &models.Writing{
		Author: &models.User{},
	}

	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&writing.ID,
		&writing.Type,
		&writing.Title,
		&writing.Content,
		&writing.AuthorID,
		&writing.CreatedAt,
		&writing.UpdatedAt,
		&writing.Author.ID,
		&writing.Author.Name,
		&writing.Author.Lastname,
		&writing.Author.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	details, err := s.FindDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	writing.Tags = details.Tags

	return writing, nil
}

func (s WritingStore) FindDetails(ctx context.Context, id string) (*models.WritingDetails, error) {
	query := `
		SELECT name 
		FROM tag_writing
		JOIN tags ON tag_writing.tag_id = tags.id
		WHERE writing_id = $1		
	`

	rows, err := s.db.QueryContext(ctx, query, id)
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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	writing := &models.WritingDetails{
		ID:   id,
		Tags: tags,
	}

	return writing, nil
}

func (s WritingStore) ComposeFeed(ctx context.Context, userID *string) ([]*models.Writing, error) {
	if userID == nil {

	}

	return make([]*models.Writing, 0), nil
}

func (s WritingStore) Update(ctx context.Context, writting *models.Writing) error {
	query := `
		UPDATE writings 
		SET type = $1, title = $2, content = $3, updated_at = NOW()
		WHERE id = $4;
	`

	args := []interface{}{writting.Type, writting.Title, writting.Content, writting.ID}

	_, err := s.db.ExecContext(ctx, query, args...)

	return err
}

func (s WritingStore) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM writings 
		WHERE id = $1
	`

	_, err := s.db.ExecContext(ctx, query, id)

	return err
}

func (s WritingStore) fetchLikesForWritings(ctx context.Context, writingsIDs []string) (map[string][]*models.WritingLike, error) {
	likesQuery := `
		SELECT writing_likes.writing_id, writing_likes.author_id, writing_likes.created_at, 
       users.id AS user_id, users.name, users.lastname
		FROM writing_likes
		LEFT JOIN users ON writing_likes.author_id = users.id
		WHERE writing_id = ANY($1)
	`

	likesRows, err := s.db.QueryContext(ctx, likesQuery, pq.Array(writingsIDs))
	if err != nil {
		return nil, err
	}
	defer likesRows.Close()

	likesMap := map[string][]*models.WritingLike{}

	for likesRows.Next() {
		like := &models.WritingLike{}
		author := &models.User{}
		var writingID string

		err := likesRows.Scan(&writingID, &like.AuthorID, &like.CreatedAt, &author.ID, &author.Name, &author.Lastname)
		if err != nil {
			return nil, err
		}

		like.Author = author
		likesMap[writingID] = append(likesMap[writingID], like)
	}

	if err = likesRows.Err(); err != nil {
		return nil, err
	}

	return likesMap, nil
}

func (s WritingStore) fetchCommentsForWritings(ctx context.Context, writingsIDs []string) (map[string][]*models.Comment, error) {
	commentsQuery := `
		SELECT comments.id, comments.content, comments.user_id, comments.created_at, 
			 users.id AS user_id, users.name, users.lastname
		FROM comments
		LEFT JOIN users ON comments.user_id = users.id
		WHERE writing_id = ANY($1)
	`

	commentsRows, err := s.db.QueryContext(ctx, commentsQuery, pq.Array(writingsIDs))
	if err != nil {
		return nil, err
	}
	defer commentsRows.Close()

	commentsMap := map[string][]*models.Comment{}

	for commentsRows.Next() {
		comment := &models.Comment{}
		author := &models.User{}
		var writingID string

		err := commentsRows.Scan(&comment.ID, &comment.Content, &comment.AuthorID, &comment.CreatedAt, &author.ID, &author.Name, &author.Lastname)
		if err != nil {
			return nil, err
		}

		comment.Author = author
		commentsMap[writingID] = append(commentsMap[writingID], comment)
	}

	if err = commentsRows.Err(); err != nil {
		return nil, err
	}

	return commentsMap, nil
}
