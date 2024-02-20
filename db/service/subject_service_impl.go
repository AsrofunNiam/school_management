package service

import (
	"errors"
	"fmt"

	"github.com/aadgraha/school_management/db/repository"
	db "github.com/aadgraha/school_management/db/sqlc"
	"github.com/go-playground/validator/v10"
)

type SubjectServiceImpl struct {
	SubJectRepository repository.SubJectRepository
	Validate          *validator.Validate
}

func NewSubjectServiceImpl(
	subjectRepository repository.SubJectRepository,
	validate *validator.Validate,
) SubjectService {
	return &SubjectServiceImpl{
		SubJectRepository: subjectRepository,
		Validate:          validate,
	}
}

func (service *SubjectServiceImpl) FindById(id string) (*db.Subject, error) {
	category := service.SubJectRepository.FindById(id)

	if category == nil {
		return category, errors.New(" Category Not Found")

	} else {
		fmt.Println(category)
		return category, nil
	}
}

func (service *SubjectServiceImpl) Create(param db.InsertSubjectParams) (*db.Subject, error) {
	// query := &db.Queries{}
	// data, err := query.InsertSubject(context.Background(), db.InsertSubjectParams(param))
	// err := service.Validate.Struct(param)

	request := &db.InsertSubjectParams{
		ID:   param.ID,
		Name: param.Name,
	}

	category := service.SubJectRepository.Create(request)

	if category == nil {
		return category, errors.New(" Category Not Found")
	} else {
		fmt.Println(category)
		return category, nil
	}
}
