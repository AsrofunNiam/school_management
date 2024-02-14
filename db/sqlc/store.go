package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v ", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateTeacherResult struct {
	User    User    `json:"user"`
	Teacher Teacher `json:"teacher"`
}

func (store *Store) CreateTeacherTx(ctx context.Context, arg CreateUserParams) (CreateTeacherResult, error) {
	var result CreateTeacherResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error
		result.User, err = q.CreateUser(ctx, arg)
		if arg.Role != "teacher" {
			err = fmt.Errorf("role is invalid")
		}
		if err != nil {
			return err
		}

		result.Teacher, err = q.CreateTeacher(ctx, arg.ID)
		if err != nil {
			return err
		}
		return nil

	})

	return result, err
}
