-- name: CreateTeacher :one
INSERT INTO teachers (
	nip,
	user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetTeacher :one
SELECT * FROM teachers
WHERE id = $1 LIMIT 1;

-- name: ListTeachers :many
SELECT * FROM teachers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTeacher :one
UPDATE teachers
SET
	nip = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTeacher :exec
DELETE FROM teachers
WHERE id = $1;