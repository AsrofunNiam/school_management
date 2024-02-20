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

func (repository *SubjectRepositoryImpl) Update(subjectRequest *db.UpdateSubjectNewParams) *db.Subject {
	query := &db.Queries{}
	data, err := query.UpdateSubjectNew(context.Background(), db.UpdateSubjectNewParams(*subjectRequest))
	helper.PanicIfError(err)

	return &data
}

func (repository *SubjectRepositoryImpl) Delete(id int64) error {
	query := &db.Queries{}
	err := query.DeleteSubject(context.Background(), id)

	return err

}
