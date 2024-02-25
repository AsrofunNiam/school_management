// All Function Service
package repository

import (
	"context"

	"github.com/aadgraha/school_management/helper"
	db "github.com/aadgraha/school_management/model/sqlc"
	"github.com/stretchr/testify/mock"
)

type SubjectRepositoryMockImpl struct {
	Mock mock.Mock
}

func NewSubjectRepositoryMockImpl() SubJectRepositoryMock {
	return &SubjectRepositoryMockImpl{}
}

func (repository *SubjectRepositoryMockImpl) FindById(id string) db.Subject {
	arguments := repository.Mock.Called(id)
	// if arguments.Get(0) == nil {
	// 	return nil

	// } else {
	category := arguments.Get(0).(db.Subject)
	return category
	// }
}

func (repository *SubjectRepositoryMockImpl) Create(subjectRequest *db.InsertSubjectParams) *db.Subject {
	query := &db.Queries{}
	data, err := query.InsertSubject(context.Background(), db.InsertSubjectParams(*subjectRequest))

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Update(subjectRequest *db.UpdateSubjectNewParams) *db.Subject {
	query := &db.Queries{}
	data, err := query.UpdateSubjectNew(context.Background(), db.UpdateSubjectNewParams{
		Name: subjectRequest.Name,
		ID:   subjectRequest.ID,
		ID_2: subjectRequest.ID_2,
	})

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Delete(id int64) error {
	query := &db.Queries{}
	err := query.DeleteSubject(context.Background(), id)

	return err

}
