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

func (repository *SubjectRepositoryImpl) FindAll(dbx *dbx.Connect) []db.Subject {
	subjectData, err := dbx.Query.SelectSubjects(context.Background())
	helper.PanicIfError(err)

	return subjectData
}

func (repository *SubjectRepositoryImpl) FindById(dbx *dbx.Connect, id string) db.Subject {
	newID, _ := strconv.ParseInt(id, 10, 64)
	subjectData, err := dbx.Query.SelectSubject(context.Background(), newID)
	helper.PanicIfError(err)

	return subjectData
}

func (repository *SubjectRepositoryImpl) Create(dbx *dbx.Connect, subjectRequest *db.InsertSubjectParams) *db.Subject {
	subjectData, err := dbx.Query.InsertSubject(context.Background(), db.InsertSubjectParams(*subjectRequest))
	helper.PanicIfError(err)

	return &subjectData
}

func (repository *SubjectRepositoryImpl) Update(dbx *dbx.Connect, id string, subjectRequest *db.UpdateSubjectNewParams) *db.Subject {
	subjectData, err := dbx.Query.UpdateSubjectNew(context.Background(), *subjectRequest)
	helper.PanicIfError(err)
	return &subjectData
}

func (repository *SubjectRepositoryImpl) Delete(dbx *dbx.Connect, id string) db.Subject {
	newID, _ := strconv.ParseInt(id, 10, 64)
	subjectData, err := dbx.Query.SelectSubject(context.Background(), newID)
	helper.PanicIfError(err)
	dbx.Query.DeleteSubject(context.Background(), newID)

	return subjectData

}
