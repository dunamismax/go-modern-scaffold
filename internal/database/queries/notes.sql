-- name: GetNote :one
SELECT * FROM notes
WHERE id = $1 LIMIT 1;

-- name: ListNotes :many
SELECT * FROM notes
ORDER BY created_at DESC;

-- name: CreateNote :one
INSERT INTO notes (
  content
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateNote :one
UPDATE notes
SET
  content = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;
