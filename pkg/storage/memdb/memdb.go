package memdb

import (
	"fmt"
	"go_news/pkg/storage"
	"time"
)

// Хранилище данных.
type Storage struct {
	topID int
	data  map[int]storage.Post
}

// Конструктор объекта хранилища.
func New() *Storage {
	var s = Storage{
		topID: 0,
		data:  make(map[int]storage.Post),
	}
	return &s
}

func (s *Storage) Posts() ([]storage.Post, error) {
	var res []storage.Post
	for _, v := range s.data {
		res = append(res, v)
	}
	return res, nil
}

func (s *Storage) AddPost(post storage.Post) error {
	_, ok := s.data[post.ID]
	if ok {
		return fmt.Errorf("post %d already in database", post.ID)
	}
	post.CreatedAt = time.Now().Unix()
	s.topID++
	post.ID = s.topID
	s.data[s.topID] = post
	return nil
}

func (s *Storage) UpdatePost(post storage.Post) error {
	_, ok := s.data[post.ID]
	if !ok {
		return fmt.Errorf("post %d is not in database", post.ID)
	}
	s.data[post.ID] = post
	return nil
}

func (s *Storage) DeletePost(post storage.Post) error {
	_, ok := s.data[post.ID]
	if !ok {
		return fmt.Errorf("post %d is not in database", post.ID)
	}
	delete(s.data, post.ID)
	return nil
}
