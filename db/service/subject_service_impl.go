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

func (service *SubjectServiceImpl) FindById(id string) (db.Subject, error) {
	category := service.SubJectRepository.FindById(id)

	// if category == nil {
	// 	return category, errors.New(" Category Not Found")

	// } else {
	fmt.Println(category)
	return category, nil
	// }
}

func (service *SubjectServiceImpl) Create(param db.InsertSubjectParams) (*db.Subject, error) {
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

func (service *SubjectServiceImpl) Update(param db.UpdateSubjectNewParams) (*db.Subject, error) {
	request := &db.UpdateSubjectNewParams{
		ID:   param.ID,
		Name: param.Name,
		ID_2: param.ID_2,
	}

	category := service.SubJectRepository.Update(request)
	if category == nil {
		return category, errors.New(" Id Not Found")
	} else {
		fmt.Println(category)
		return category, nil
	}
}

func (service *SubjectServiceImpl) Delete(id int64) error {
	service.SubJectRepository.Delete(id)

	return errors.ErrUnsupported
}
