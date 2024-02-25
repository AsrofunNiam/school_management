-- name: CreateTeacher :one
INSERT INTO teachers (
	nip
) VALUES (
  $1
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
