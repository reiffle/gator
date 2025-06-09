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


-- name: MarkFeedsFetched :exec
UPDATE feeds
SET updated_at = $1, last_fetched_at = $1
WHERE id = $2;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
INNER JOIN feed_follows on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT 1;