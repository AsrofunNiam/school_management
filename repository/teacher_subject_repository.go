package repository

import (
	dbx "github.com/aadgraha/school_management/model"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type TeacherSubjectRepository interface {
	Create(dbx *dbx.Connect, subject *db.InsertTeacherSubjectParams) *db.TeacherSubject
	FindAll(dbx *dbx.Connect) []db.SelectTeacherSubjectsRow
	FindById(dbx *dbx.Connect, id string) db.SelectTeacherSubjectRow
	Update(dbx *dbx.Connect, id string, subject *db.UpdateTeacherSubjectNewParams) *db.TeacherSubject
	Delete(dbx *dbx.Connect, id string) db.SelectTeacherSubjectRow
}
