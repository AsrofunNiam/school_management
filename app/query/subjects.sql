-- name: SelectSubjects :many
SELECT *
FROM subjects 
;

-- name: SelectSubject :one
SELECT *
FROM subjects 
WHERE id = $1;


-- name: DeleteSubject :exec
DELETE
FROM subjects 
WHERE id = $1
RETURNING id, name; 

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

-- name: UpdateSubjectNew :one
UPDATE subjects
SET id = $1,
    name = $2
WHERE id = $3
RETURNING *;
