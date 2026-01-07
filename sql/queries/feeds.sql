-- name: CreateFeed :one
INSERT INTO
    feeds (name, url, user_id)
VALUES
    ($1,$2,$3)
RETURNING
    *
;

-- name: GetFeeds :many
SELECT
    *
FROM
    feeds
;

-- name: GetFeedsWithUserName :many
SELECT
    feeds.id,
    feeds.name,
    feeds.url,
    feeds.user_id,
    feeds.created_at,
    feeds.updated_at,
    users.name AS username
FROM
    feeds
INNER JOIN users
    ON feeds.user_id = users.id
;

-- name: GetFeedByUrl :one
SELECT
    *
FROM
    feeds
WHERE 
    url = $1
;