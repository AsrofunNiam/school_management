package repository

import db "github.com/aadgraha/school_management/model/sqlc"

type SubJectRepository interface {
	FindById(id string) db.Subject
	Create(subject *db.InsertSubjectParams) *db.Subject
	Update(subject *db.UpdateSubjectNewParams) *db.Subject
	Delete(id int64) error

	// testing
	// FindLengthAll() []db.Subject
}
