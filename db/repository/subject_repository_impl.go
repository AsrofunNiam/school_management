package repository

import (
	"context"

	"github.com/aadgraha/school_management/db/helper"
	db "github.com/aadgraha/school_management/db/sqlc"
)

type SubjectRepositoryImpl struct {
}

func NewSubjectRepository() SubJectRepository {
	return &SubjectRepositoryImpl{}
}

func (repository *SubjectRepositoryImpl) FindById(id string) *db.Subject {
	data := db.Subject{}
	return &data
}

func (repository *SubjectRepositoryImpl) Create(subjectRequest *db.InsertSubjectParams) *db.Subject {
	query := &db.Queries{}
	data, err := query.InsertSubject(context.Background(), db.InsertSubjectParams(*subjectRequest))
	helper.PanicIfError(err)

	return &data
}
