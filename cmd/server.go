package main

import (
	"go_news/pkg/api"
	"go_news/pkg/storage"
	"go_news/pkg/storage/mongodb"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	var srv server

	// db := memdb.New()

	/*
		// Реляционная БД PostgreSQL.
		db2, err := postgres.New("postgresql://localhost/devbase?user=postgres&password=postgres")
		if err != nil {
			log.Fatal(err)
		}
	*/

	// Документная БД MongoDB.
	db3, err := mongodb.New("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal(err)
	}

	srv.db = db3

	srv.api = api.New(srv.db)

	http.ListenAndServe(":8080", srv.api.Router())
}
