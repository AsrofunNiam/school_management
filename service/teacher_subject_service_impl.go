package service

import (
	"database/sql"
	"errors"
	"fmt"

	dbx "github.com/aadgraha/school_management/model"
	teacher_subject "github.com/aadgraha/school_management/model/sqlc"
	"github.com/aadgraha/school_management/repository"
	"github.com/go-playground/validator/v10"
)

type TeacherSubjectServiceImpl struct {
	TeacherSubjectRepository repository.TeacherSubjectRepository
	DB                       *dbx.Connect
	Validate                 *validator.Validate
}

func NewTeacherSubjectServiceImpl(
	TeacherSubjectRepository repository.TeacherSubjectRepository,
	db *sql.DB,
	dbx *dbx.Connect,
	validate *validator.Validate,
) TeacherSubjectService {
	return &TeacherSubjectServiceImpl{
		TeacherSubjectRepository: TeacherSubjectRepository,
		Validate:                 validate,
		DB:                       dbx,
	}
}

func (service *TeacherSubjectServiceImpl) FindAll() ([]teacher_subject.SelectTeacherSubjectsRow, error) {
	tx := service.DB
	request := service.TeacherSubjectRepository.FindAll(tx)

	if tx == nil {
		return request, errors.New(" Failed call db")

	} else {
		return request, nil
	}
}

func (service *TeacherSubjectServiceImpl) FindById(id string) (teacher_subject.SelectTeacherSubjectRow, error) {
	tx := service.DB
	request := service.TeacherSubjectRepository.FindById(tx, id)

	if tx == nil {
		return request, errors.New(" Failed call db")

	} else {
		return request, nil
	}
}

func (service *TeacherSubjectServiceImpl) Create(param teacher_subject.InsertTeacherSubjectParams) (*teacher_subject.TeacherSubject, error) {
	tx := service.DB
	request := &teacher_subject.InsertTeacherSubjectParams{
		ID:        param.ID,
		TeacherID: param.TeacherID,
		SubjectID: param.SubjectID,
		Period:    param.Period,
	}

	category := service.TeacherSubjectRepository.Create(tx, request)

	if category == nil {
		return category, errors.New("Create Failed")
	} else {
		return category, nil
	}
}

func (service *TeacherSubjectServiceImpl) Update(id string, param teacher_subject.UpdateTeacherSubjectNewParams) (*teacher_subject.TeacherSubject, error) {
	tx := service.DB
	request := &teacher_subject.UpdateTeacherSubjectNewParams{
		ID:        param.ID,
		TeacherID: param.TeacherID,
		SubjectID: param.SubjectID,
		Period:    param.Period,
		// ID_2: param.ID_2,
	}

	category := service.TeacherSubjectRepository.Update(tx, id, request)
	if category == nil {
		return category, errors.New(" Id Not Found")
	} else {
		fmt.Println(category)
		return category, nil
	}
}

func (service *TeacherSubjectServiceImpl) Delete(id string) (teacher_subject.SelectTeacherSubjectRow, error) {
	tx := service.DB
	selectCategory := service.TeacherSubjectRepository.FindById(tx, id)
	service.TeacherSubjectRepository.Delete(tx, id)

	return selectCategory, nil
}
