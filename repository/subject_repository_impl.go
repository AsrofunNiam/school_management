package repository

import (
	"context"
	"strconv"

	"github.com/aadgraha/school_management/helper"
	db "github.com/aadgraha/school_management/model/sqlc"
)

type SubjectRepositoryImpl struct {
}

func NewSubjectRepository() SubJectRepository {
	return &SubjectRepositoryImpl{}
}

func (repository *SubjectRepositoryImpl) FindById(id string) db.Subject {
	newID, _ := strconv.ParseInt(id, 10, 64)
	// data := db.Subject{}
	var query *db.Queries
	subject, err := query.SelectSubject(context.Background(), newID)
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

// func (repository *SubjectRepositoryImpl) FindLengthAll() []db.Subject {
// 	query := &db.Queries{}
// 	test := query.SelectSubject(context.Background())
// 	// err := query.DeleteSubject(context.Background(), id)

// 	return err

// }
