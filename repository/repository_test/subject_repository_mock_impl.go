// All Function Service
package repository

import (
	"context"

	"github.com/aadgraha/school_management/helper"
	dbx "github.com/aadgraha/school_management/model"
	subject "github.com/aadgraha/school_management/model/sqlc"
	"github.com/stretchr/testify/mock"
)

type SubjectRepositoryMockImpl struct {
	Mock mock.Mock
	DB   *dbx.Connect
}

func NewSubjectRepositoryMockImpl() SubJectRepositoryMock {
	return &SubjectRepositoryMockImpl{}
}

func (repository *SubjectRepositoryMockImpl) FindById(dbx *dbx.Connect, id string) subject.Subject {
	arguments := repository.Mock.Called(id)
	// if arguments.Get(0) == nil {
	// 	return nil

	// } else {
	category := arguments.Get(0).(subject.Subject)
	return category
	// }
}

func (repository *SubjectRepositoryMockImpl) Create(dbx *dbx.Connect, subjectRequest *subject.InsertSubjectParams) *subject.Subject {
	query := &subject.Queries{}
	data, err := query.InsertSubject(context.Background(), subject.InsertSubjectParams(*subjectRequest))

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Update(dbx *dbx.Connect, subjectRequest *subject.UpdateSubjectNewParams) *subject.Subject {
	query := &subject.Queries{}
	data, err := query.UpdateSubjectNew(context.Background(), subject.UpdateSubjectNewParams{
		Name: subjectRequest.Name,
		ID:   subjectRequest.ID,
		ID_2: subjectRequest.ID_2,
	})

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Delete(dbx *dbx.Connect, id int64) error {
	query := &subject.Queries{}
	err := query.DeleteSubject(context.Background(), id)

	return err

}
