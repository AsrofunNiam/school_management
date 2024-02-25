package repository

import (
	"context"
	"strconv"

	"github.com/aadgraha/school_management/helper"
	dbx "github.com/aadgraha/school_management/model"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type SubjectRepositoryImpl struct {
}

func NewSubjectRepository() SubJectRepository {
	return &SubjectRepositoryImpl{}
}

func (repository *SubjectRepositoryImpl) FindById(dbx *dbx.Connect, id string) db.Subject {
	newID, _ := strconv.ParseInt(id, 10, 64)
	subject, err := dbx.Query.SelectSubject(context.Background(), newID)
	helper.PanicIfError(err)

	return subject
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
