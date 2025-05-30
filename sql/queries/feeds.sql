-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: PrintFeeds :many

SELECT feeds.name, feeds.url, users.name
FROM feeds
INNER JOIN users on users.id = feeds.user_id;