package service

import (
	"database/sql"
	"errors"
	"fmt"

	dbx "github.com/aadgraha/school_management/model"
	subject "github.com/aadgraha/school_management/model/sqlc"
	"github.com/aadgraha/school_management/repository"
	"github.com/go-playground/validator/v10"
)

type SubjectServiceImpl struct {
	SubJectRepository repository.SubJectRepository
	DB                *dbx.Connect
	Validate          *validator.Validate
}

func NewSubjectServiceImpl(
	subjectRepository repository.SubJectRepository,
	db *sql.DB,
	dbx *dbx.Connect,
	validate *validator.Validate,
) SubjectService {
	return &SubjectServiceImpl{
		SubJectRepository: subjectRepository,
		Validate:          validate,
		DB:                dbx,
	}
}

func (service *SubjectServiceImpl) FindById(id string) (subject.Subject, error) {
	tx := service.DB
	category := service.SubJectRepository.FindById(tx, id)

	if tx == nil {
		return category, errors.New(" Failed call db")

	} else {
		return category, nil
	}
}

func (service *SubjectServiceImpl) Create(param subject.InsertSubjectParams) (*subject.Subject, error) {
	tx := service.DB
	request := &subject.InsertSubjectParams{
		ID:   param.ID,
		Name: param.Name,
	}

	category := service.SubJectRepository.Create(tx, request)

	if category == nil {
		return category, errors.New("Create Failed")
	} else {
		return category, nil
	}
}

func (service *SubjectServiceImpl) Update(id string, param subject.UpdateSubjectNewParams) (*subject.Subject, error) {
	tx := service.DB
	request := &subject.UpdateSubjectNewParams{
		ID:   param.ID,
		Name: param.Name,
		ID_2: param.ID_2,
	}

	category := service.SubJectRepository.Update(tx, id, request)
	if category == nil {
		return category, errors.New(" Id Not Found")
	} else {
		fmt.Println(category)
		return category, nil
	}
}

func (service *SubjectServiceImpl) Delete(id string) (subject.Subject, error) {
	tx := service.DB
	selectCategory := service.SubJectRepository.FindById(tx, id)
	service.SubJectRepository.Delete(tx, id)

	return selectCategory, nil
}
