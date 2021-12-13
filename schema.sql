DROP SCHEMA IF EXISTS news;

CREATE SCHEMA IF NOT EXISTS news
    AUTHORIZATION postgres;

GRANT ALL ON SCHEMA news TO developer;

GRANT ALL ON SCHEMA news TO postgres;

ALTER DEFAULT PRIVILEGES IN SCHEMA news
GRANT ALL ON TABLES TO developer;

DROP TABLE IF EXISTS
    devbase.news.authors,
    devbase.news.posts;

CREATE TABLE devbase.news.authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE devbase.news.posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES devbase.news.authors(id),
    created_at BIGINT NOT NULL,
    published_at BIGINT DEFAULT 0,
    title TEXT,
    content TEXT
);