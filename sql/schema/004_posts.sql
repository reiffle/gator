-- +goose up
CREATE TABLE posts (
    id              UUID PRIMARY KEY,
    created_at      TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    title           TEXT NOT NULL,
    url             TEXT NOT NULL,
    description     TEXT,
    published_at    TIMESTAMP,
    feed_id         UUID NOT NULL references feeds(id) on DELETE cascade
);

-- +goose down
DROP TABLE posts;