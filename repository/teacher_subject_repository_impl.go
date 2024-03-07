package repository

import (
	"context"
	"strconv"

	"github.com/aadgraha/school_management/helper"
	dbx "github.com/aadgraha/school_management/model"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type TeacherSubjectRepositoryImpl struct {
}

func NewTeacherSubjectRepository() TeacherSubjectRepository {
	return &TeacherSubjectRepositoryImpl{}
}

func (repository *TeacherSubjectRepositoryImpl) Create(dbx *dbx.Connect, subjectRequest *db.InsertTeacherSubjectParams) *db.TeacherSubject {
	subjectData, err := dbx.Query.InsertTeacherSubject(context.Background(), db.InsertTeacherSubjectParams(*subjectRequest))
	helper.PanicIfError(err)

	return &subjectData
}

func (repository *TeacherSubjectRepositoryImpl) FindAll(dbx *dbx.Connect) []db.SelectTeacherSubjectsRow {
	subjectData, err := dbx.Query.SelectTeacherSubjects(context.Background())
	helper.PanicIfError(err)

	return subjectData
}

func (repository *TeacherSubjectRepositoryImpl) FindById(dbx *dbx.Connect, id string) db.SelectTeacherSubjectRow {
	newID, _ := strconv.ParseInt(id, 10, 64)
	subjectData, err := dbx.Query.SelectTeacherSubject(context.Background(), newID)
	helper.PanicIfError(err)

	return subjectData
}

func (repository *TeacherSubjectRepositoryImpl) Update(dbx *dbx.Connect, id string, subjectRequest *db.UpdateTeacherSubjectNewParams) *db.TeacherSubject {
	subjectData, err := dbx.Query.UpdateTeacherSubjectNew(context.Background(), *subjectRequest)
	helper.PanicIfError(err)
	return &subjectData
}

func (repository *TeacherSubjectRepositoryImpl) Delete(dbx *dbx.Connect, id string) db.SelectTeacherSubjectRow {
	newID, _ := strconv.ParseInt(id, 10, 64)
	subjectData, err := dbx.Query.SelectTeacherSubject(context.Background(), newID)
	helper.PanicIfError(err)
	dbx.Query.DeleteSubject(context.Background(), newID)

	return subjectData

}
