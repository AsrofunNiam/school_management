package service

import (
	db "github.com/aadgraha/school_management/model/sqlc"
)

type TeacherSubjectService interface {
	FindAll() ([]db.SelectTeacherSubjectsRow, error)
	FindById(id string) (db.SelectTeacherSubjectRow, error)
	Create(param db.InsertTeacherSubjectParams) (*db.TeacherSubject, error)
	Update(id string, param db.UpdateTeacherSubjectNewParams) (*db.TeacherSubject, error)
	Delete(id string) (db.SelectTeacherSubjectRow, error)

	// test length
	// FindLengthAll() []db.Subject
}
