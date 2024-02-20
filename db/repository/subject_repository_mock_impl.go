// All Function Service
package repository

import (
	"context"

	"github.com/aadgraha/school_management/db/helper"
	db "github.com/aadgraha/school_management/db/sqlc"
	"github.com/stretchr/testify/mock"
)

type SubjectRepositoryMockImpl struct {
	Mock mock.Mock
}

func NewSubjectRepositoryMockImpl() SubJectRepositoryMock {
	return &SubjectRepositoryMockImpl{}
}

func (repository *SubjectRepositoryMockImpl) FindById(id string) *db.Subject {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil

	} else {
		category := arguments.Get(0).(db.Subject)
		return &category
	}
}

func (repository *SubjectRepositoryMockImpl) Create(subjectRequest *db.InsertSubjectParams) *db.Subject {
	query := &db.Queries{}
	data, err := query.InsertSubject(context.Background(), db.InsertSubjectParams(*subjectRequest))

	helper.PanicIfError(err)
	return &data
	// data := db.Subject{}
	// return &data
	// arguments := repository.Mock.Called(id)
	// if arguments.Get(0) == nil {
	// 	return nil

	// } else {
	// 	category := arguments.Get(0).(db.Subject)
	// 	return &category
	// }
}
