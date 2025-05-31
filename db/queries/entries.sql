-- name: CreateEntry :one
INSERT INTO entries (
  id_account,
  amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id_entries = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id_entries
LIMIT $1 OFFSET $2;

-- name: ListEntriesByAccount :many
SELECT * FROM entries
WHERE id_account = $3
ORDER BY id_entries
LIMIT $1 OFFSET $2;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id_entries = $1;