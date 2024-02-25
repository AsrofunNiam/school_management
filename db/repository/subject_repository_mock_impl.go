// All Function Service
package repository

import (
	"context"

	"github.com/aadgraha/school_management/db/helper"
	subject "github.com/aadgraha/school_management/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type SubjectRepositoryMockImpl struct {
	Mock mock.Mock
}

func NewSubjectRepositoryMockImpl() SubJectRepositoryMock {
	return &SubjectRepositoryMockImpl{}
}

func (repository *SubjectRepositoryMockImpl) FindById(id string) subject.Subject {
	arguments := repository.Mock.Called(id)
	// if arguments.Get(0) == nil {
	// 	return

	// } else {
	category := arguments.Get(0).(subject.Subject)
	return category
	// }
}

func (repository *SubjectRepositoryMockImpl) Create(subjectRequest *subject.InsertSubjectParams) *subject.Subject {
	query := &subject.Queries{}
	data, err := query.InsertSubject(context.Background(), subject.InsertSubjectParams(*subjectRequest))

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Update(subjectRequest *subject.UpdateSubjectNewParams) *subject.Subject {
	query := &subject.Queries{}
	data, err := query.UpdateSubjectNew(context.Background(), subject.UpdateSubjectNewParams{
		Name: subjectRequest.Name,
		ID:   subjectRequest.ID,
		ID_2: subjectRequest.ID_2,
	})

	helper.PanicIfError(err)
	return &data
}

func (repository *SubjectRepositoryMockImpl) Delete(id int64) error {
	query := &subject.Queries{}
	err := query.DeleteSubject(context.Background(), id)

	return err

}
