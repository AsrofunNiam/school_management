-- name: CreateUser :one
INSERT INTO users (
	id,
	name,
	role
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT id, name, role FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, name, role FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
	id = $4,
	name = $2,
	role = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;