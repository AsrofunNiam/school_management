package repository

import (
	dbx "github.com/aadgraha/school_management/model"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type SubJectRepository interface {
	FindAll(dbx *dbx.Connect) []db.Subject
	FindById(dbx *dbx.Connect, id string) db.Subject
	Create(dbx *dbx.Connect, subject *db.InsertSubjectParams) *db.Subject
	Update(dbx *dbx.Connect, id string, subject *db.UpdateSubjectNewParams) *db.Subject
	Delete(dbx *dbx.Connect, id string) db.Subject

	// testing
	// FindLengthAll() []db.Subject
}
