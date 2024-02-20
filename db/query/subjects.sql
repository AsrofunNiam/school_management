-- name: SelectSubjects :many
SELECT *
FROM subjects 
;

-- name: SelectSubject :one
SELECT *
FROM subjects 
WHERE id = $1;


-- name: DeleteSubject :exec
SELECT *
FROM subjects 
WHERE id = $1; 

-- name: InsertSubject :one
INSERT INTO subjects (
    id,
    name
    )VALUES (
    $1,
    $2
    )
RETURNING *;

-- name: UpdateSubject :exec
UPDATE subjects
SET name = coalesce(sqlc.narg('name'), slug)
WHERE id = sqlc.arg('id')
RETURNING *;
