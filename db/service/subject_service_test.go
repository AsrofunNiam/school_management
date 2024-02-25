package service

import (
	"fmt"
	"testing"

	"github.com/aadgraha/school_management/db/repository"
	subject "github.com/aadgraha/school_management/db/sqlc"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// set var call data mock
var categoryRepositoryMock = repository.SubjectRepositoryMockImpl{
	Mock: mock.Mock{},
}
var categoryServices = SubjectServiceImpl{
	SubJectRepository: &categoryRepositoryMock,
	Validate:          &validator.Validate{},
}

func TestSubjectServices_GetNotFound(t *testing.T) {

	//  implement calling mock
	categoryRepositoryMock.Mock.On("FindById", "1").Return(nil)

	// send param to repository
	category, err := categoryServices.FindById("1")

	fmt.Println(category)
	fmt.Println(err)

	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestSubjectServices_GetSuccess(t *testing.T) {

	// Return  value pointer
	category := subject.Subject{
		ID:   3,
		Name: "Budi",
	}

	// Call by param
	categoryRepositoryMock.Mock.On("FindById", "2").Return(category)

	result, err := categoryServices.FindById("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)

}

func TestSubjectServices_Create(t *testing.T) {
	// Return  value pointer
	category := &subject.InsertSubjectParams{
		ID:   3,
		Name: "Budi",
	}

	// Call by param
	categoryRepositoryMock.Mock.On("Create", category).Return(category)

	result, err := categoryServices.Create(*category)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)

}

func TestSubjectServices_UpdateSuccess(t *testing.T) {
	// Return  value pointer
	response := &subject.UpdateSubjectNewParams{
		ID:   3,
		Name: "Budi",
		ID_2: 2,
	}

	// Call by param
	categoryRepositoryMock.Mock.On("Update", response.ID_2).Return(response)

	result, err := categoryServices.Update(*response)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, response.ID, result.ID)
	assert.Equal(t, response.Name, result.Name)

}

func TestSubjectServices_Delete(t *testing.T) {
	// Return  value pointer
	category := &subject.InsertSubjectParams{
		ID:   2,
		Name: "Budi",
	}

	// Call by param
	categoryRepositoryMock.Mock.On("Delete", 2).Return(category)

	result := categoryServices.Delete(2)
	assert.NotNil(t, result)

}
