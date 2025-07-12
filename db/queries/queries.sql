-- name: GetMessages :many
SELECT * FROM messages ORDER BY created_at DESC;

-- name: CreateMessage :one
INSERT INTO messages (body) VALUES ($1) RETURNING *;
