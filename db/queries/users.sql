-- name: CreateUser :one
INSERT INTO Users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM Users
WHERE username = $1 LIMIT 1;

