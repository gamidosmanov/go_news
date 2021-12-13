package postgres

import (
	"context"
	"go_news/pkg/storage"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Database abstraction
type Storage struct {
	db *pgxpool.Pool
}

// New returns Storage object with pg connection pool
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{db: db}
	return &s, nil
}

func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			posts.id,
			posts.title,
			posts.content,
			posts.author_id,
			authors.name,
			posts.created_at,
			posts.published_at
		FROM
			devbase.news.posts AS posts,
			devbase.news.authors AS authors
		WHERE
			posts.author_id = authors.id
	`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.CreatedAt,
			&p.PublishedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (s *Storage) AddPost(p storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO devbase.news.posts (author_id, created_at, title, content)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`,
		p.AuthorID,
		time.Now().Unix(),
		p.Title,
		p.Content,
	).Scan(&id)
	return err
}

func (s *Storage) UpdatePost(p storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		UPDATE devbase.news.posts
		SET
			author_id = $2,
			published_at = $3,
			title = $4,
			content = $5
		WHERE id = $1 RETURNING id
		`,
		p.ID,
		p.AuthorID,
		p.PublishedAt,
		p.Title,
		p.Content,
	).Scan(&id)
	return err
}

func (s *Storage) DeletePost(p storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		DELETE FROM devbase.news.posts
		WHERE id = $1 RETURNING id
		`,
		p.ID,
	).Scan(&id)
	return err
}
