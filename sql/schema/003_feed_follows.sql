-- +goose up

CREATE TABLE feed_follows (
    id          UUID PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    user_id     UUID NOT NULL references users(id) on DELETE cascade,
    feed_id     UUID NOT NULL references feeds(id) on DELETE cascade,
    unique (user_id, feed_id)
);

-- +goose down
DROP TABLE feed_follows;