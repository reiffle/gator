-- +goose Up
CREATE TABLE feeds (
id              UUID PRIMARY KEY,
created_at      TIMESTAMP NOT NULL,
updated_at      TIMESTAMP NOT NULL,
name            TEXT NOT NULL,
url             TEXT UNIQUE NOT NULL,
user_id         UUID NOT NULL references users(id) on DELETE cascade,
last_fetched_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
DROP TABLE feeds;