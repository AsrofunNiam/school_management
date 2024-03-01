-- name: InsertTeacherSubject :one
INSERT INTO teacher_subjects (
    id,
    teacher_id,
    subject_id,
    period
    )VALUES (
    $1,
    $2,
    $3,
    $4
    )
RETURNING *;

-- name: SelectTeacherSubjects :many
SELECT *
FROM teacher_subjects
;

-- name: SelectTeacherSubject :one
SELECT *
FROM teacher_subjects
WHERE id = $1 
;

-- name: SelectJointTeacherSubjects :many
SELECT t1.id, t1.subject_id, t1.teacher_id, t1.period, coalesce(t2.name, 'Register Subject')  As name_subject
FROM teacher_subjects t1
LEFT JOIN subjects t2 ON t2.id = t1.subject_id 
;


-- name: SelectJointTeacherSubject :one
SELECT t1.id, t1.subject_id, t1.teacher_id, t1.period, coalesce(t2.name, 'Register Subject')  As name_subject
FROM teacher_subjects t1
LEFT JOIN subjects t2 ON t2.id = t1.subject_id
WHERE t1.id = $1;

-- name: UpdateTeacherSubject :exec
UPDATE teacher_subjects
SET subject_id = coalesce(sqlc.narg('subject_id'), slug),
    teacher_id =  coalesce(sqlc.narg('teacher_id'), slug),
    period =  coalesce(sqlc.narg('period'), slug)
    
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpdateTeacherSubjectNew :one
UPDATE teacher_subjects
SET id = $1,
    teacher_id = $2,
    subject_id = $3,
    period = $4
WHERE id = $5
RETURNING *;

-- name: DeleteTeacherSubject :exec
DELETE
FROM teacher_subjects
WHERE id = $1
RETURNING id, teacher_id, subject_id, period;