package controller

import (
	"fmt"
	"net/http"

	"github.com/aadgraha/school_management/helper"
	teacher_subject "github.com/aadgraha/school_management/model/sqlc"
	"github.com/aadgraha/school_management/model/web"
	"github.com/aadgraha/school_management/service"
	"github.com/gin-gonic/gin"
)

type TeacherSubjectControllerImpl struct {
	TeacherSubjectService service.TeacherSubjectService
}

func NewTeacherSubjectController(TeacherSubjectService service.TeacherSubjectService) TeacherSubjectController {
	return &TeacherSubjectControllerImpl{
		TeacherSubjectService: TeacherSubjectService,
	}
}

func (controller *TeacherSubjectControllerImpl) FindAll(c *gin.Context) {
	subjectResponses, err := controller.TeacherSubjectService.FindAll()
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "find teacher subjects success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TeacherSubjectControllerImpl) FindById(c *gin.Context) {
	paramID := c.Param("teacherSubjectId")
	subjectResponses, err := controller.TeacherSubjectService.FindById(paramID)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "find teacher subjects by id success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TeacherSubjectControllerImpl) Create(c *gin.Context) {
	data_request := teacher_subject.InsertTeacherSubjectParams{}
	fmt.Println("c, &request", c, &data_request)
	helper.ReadFromRequestBody(c, &data_request)
	subjectResponses, err := controller.TeacherSubjectService.Create(data_request)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "create new teacher subjects success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TeacherSubjectControllerImpl) Update(c *gin.Context) {
	paramID := c.Param("subjectId")
	data_request := teacher_subject.UpdateTeacherSubjectNewParams{}
	fmt.Println("c, &request", c, &data_request)
	helper.ReadFromRequestBody(c, &data_request)
	subjectResponses, err := controller.TeacherSubjectService.Update(paramID, data_request)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "edit teacher subjects success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)

}

func (controller *TeacherSubjectControllerImpl) Delete(c *gin.Context) {
	// paramID := c.Param("subjectId")
	// subjectResponses, err := controller.TeacherSubjectService.Delete(paramID)
	// if err != nil {
	// 	return
	// }

	webResponse := web.WebResponse{
		Success: true,
		// Message: fmt.Sprintf("delete teacher_subject %s  success", subjectResponses.TeacherID),
		Data: nil,
	}

	c.JSON(http.StatusOK, webResponse)
}
