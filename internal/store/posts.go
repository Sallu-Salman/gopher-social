package store

import (
	"context"
	"database/sql"
	"fmt"
)

type Post struct {
	Id        int    `json:"id"`
	Content   string `json:"content" validate:"required,max=100"`
	Title     string `json:"title" validate:"required,max=100"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func (p Post) String() string {
	return fmt.Sprintf(
		"Post[id: %v, title: %v, content: %v, user_id: %v, created_at: %v]",
		p.Id,
		p.Title,
		p.Content,
		p.UserId,
		p.CreatedAt,
	)
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

func (p *PostsStore) GetById(id int) (*Post, error) {
	query := `
	SELECT id, title, content, user_id, created_at
	FROM posts 
	WHERE id = $1;
	`

	post := &Post{}

	err := p.db.QueryRowContext(
		context.Background(),
		query,
		id,
	).Scan(
		&post.Id,
		&post.Title,
		&post.Content,
		&post.UserId,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}
