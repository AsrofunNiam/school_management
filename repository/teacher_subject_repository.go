package repository

import (
	dbx "github.com/aadgraha/school_management/model"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type TeacherSubjectRepository interface {
	Create(dbx *dbx.Connect, subject *db.InsertTeacherSubjectParams) *db.TeacherSubject
	FindAll(dbx *dbx.Connect) []db.TeacherSubject
	FindById(dbx *dbx.Connect, id string) db.TeacherSubject
	Update(dbx *dbx.Connect, id string, subject *db.UpdateTeacherSubjectNewParams) *db.TeacherSubject
	Delete(dbx *dbx.Connect, id string) db.TeacherSubject
}
