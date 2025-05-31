-- name: CreateTransfer :one
INSERT INTO transfers (
  from_id_account,
  to_id_account,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id_transfer = $1 LIMIT 1;


-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id_transfer
LIMIT $1 OFFSET $2;


-- name: ListTransfersByToAccount :many
SELECT * FROM transfers
WHERE to_id_account = $3
ORDER BY id_transfer
LIMIT $1 OFFSET $2;


-- name: ListTransfersByFromAccount :many
SELECT * FROM transfers
WHERE from_id_account = $3
ORDER BY id_transfer
LIMIT $1 OFFSET $2;
