package repository

import db "github.com/aadgraha/school_management/db/sqlc"

type SubJectRepository interface {
	FindById(id string) *db.Subject
	Create(subject *db.InsertSubjectParams) *db.Subject
}
