package service

import (
	db "github.com/aadgraha/school_management/db/sqlc"
)

type SubjectService interface {
	FindById(id string) (*db.Subject, error)
	Create(param db.InsertSubjectParams) (*db.Subject, error)
}
