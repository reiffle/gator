-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
        VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    returning *
)
SELECT
    inserted_feed_follow.*,
    users.name as user_name,
    feeds.name as feed_name
FROM inserted_feed_follow
    INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
    INNER JOIN users ON users.id = inserted_feed_follow.user_id;


-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*, 
    users.name as user_name,
    feeds.name as feed_name
FROM feed_follows
INNER JOIN feeds on feeds.id = feed_follows.feed_id
INNER JOIN users on users.id = feed_follows.user_id
WHERE users.name = $1;

-- name: FindFeed :one
SELECT *
FROM feeds
WHERE feeds.url = $1;