-- name: GetMessages :many
SELECT * FROM messages ORDER BY created_at DESC;

-- name: CreateMessage :exec
INSERT INTO messages (body) VALUES (?);
