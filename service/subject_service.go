package service

import (
	db "github.com/aadgraha/school_management/model/sqlc"
)

type SubjectService interface {
	FindAll() ([]db.Subject, error)
	FindById(id string) (db.Subject, error)
	Create(param db.InsertSubjectParams) (*db.Subject, error)
	Update(id string, param db.UpdateSubjectNewParams) (*db.Subject, error)
	Delete(id string) (db.Subject, error)

	// test length
	// FindLengthAll() []db.Subject
}
