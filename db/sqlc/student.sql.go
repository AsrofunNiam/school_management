// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: student.sql

package db

import (
	"context"
)

const createStudent = `-- name: CreateStudent :one
INSERT INTO students (
	nis,
	user_id
) VALUES (
  $1, $2
)
RETURNING id, nis, user_id, created_at
`

type CreateStudentParams struct {
	Nis    string `json:"nis"`
	UserID int64  `json:"user_id"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, createStudent, arg.Nis, arg.UserID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Nis,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteStudent = `-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1
`

func (q *Queries) DeleteStudent(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteStudent, id)
	return err
}

const getStudent = `-- name: GetStudent :one
SELECT id, nis, user_id, created_at FROM students
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudent(ctx context.Context, id int64) (Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, id)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Nis,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listStudents = `-- name: ListStudents :many
SELECT id, nis, user_id, created_at FROM students
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStudentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStudents(ctx context.Context, arg ListStudentsParams) ([]Student, error) {
	rows, err := q.db.QueryContext(ctx, listStudents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Student
	for rows.Next() {
		var i Student
		if err := rows.Scan(
			&i.ID,
			&i.Nis,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStudent = `-- name: UpdateStudent :one
UPDATE students
SET
	nis = $1
WHERE id = $2
RETURNING id, nis, user_id, created_at
`

type UpdateStudentParams struct {
	Nis string `json:"nis"`
	ID  int64  `json:"id"`
}

func (q *Queries) UpdateStudent(ctx context.Context, arg UpdateStudentParams) (Student, error) {
	row := q.db.QueryRowContext(ctx, updateStudent, arg.Nis, arg.ID)
	var i Student
	err := row.Scan(
		&i.ID,
		&i.Nis,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
