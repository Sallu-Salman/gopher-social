package store

import (
	"context"
	"database/sql"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	UserId int `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (p *PostsStore) Create(post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id)
	VALUES ($1, $2, $3) RETURNING id, created_at
	`
	err := p.db.QueryRowContext(
		context.Background(),
		query,
		post.Content,
		post.Title,
		post.UserId,
	).Scan(
		&post.Id,
		&post.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}