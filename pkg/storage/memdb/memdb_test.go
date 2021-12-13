package memdb

import (
	"go_news/pkg/storage"
	"reflect"
	"testing"
)

func TestStorage(t *testing.T) {
	var posts = []storage.Post{
		{
			ID:          1,
			Title:       "Effective Go",
			Content:     "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
			AuthorID:    1,
			AuthorName:  "Gamid",
			PublishedAt: 0,
		},
		{
			ID:          2,
			Title:       "The Go Memory Model",
			Content:     "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
			AuthorID:    1,
			AuthorName:  "Gamid",
			PublishedAt: 0,
		},
	}
	s := New()
	err := s.AddPost(posts[0])
	if err != nil {
		t.Error(err)
		return
	}
	res, err := s.Posts()
	if err != nil {
		t.Error(err)
		return
	}
	want := posts[0]
	got := res[0]
	got.CreatedAt = 0
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %+v\nGot %+v\n", want, got)
		return
	}
	posts[0].Content = "Hi"
	err = s.UpdatePost(posts[0])
	if err != nil {
		t.Error(err)
		return
	}
	res, err = s.Posts()
	if err != nil {
		t.Error(err)
		return
	}
	want = posts[0]
	got = res[0]
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %+v\nGot %+v\n", want, got)
		return
	}
	s.DeletePost(posts[0])
	res, err = s.Posts()
	if err != nil {
		t.Error(err)
		return
	}
	var want2 []storage.Post
	got2 := res
	if !reflect.DeepEqual(want2, got2) {
		t.Errorf("Want %+v\nGot %+v\n", want, got)
		return
	}
	err = s.UpdatePost(posts[1])
	if err != nil {
		if err.Error() != "post 2 is not in database" {
			t.Error("Incorrect update")
		}
	}
}
