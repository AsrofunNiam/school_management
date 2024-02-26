package controller

import (
	"fmt"
	"net/http"

	"github.com/aadgraha/school_management/helper"
	subject "github.com/aadgraha/school_management/model/sqlc"
	"github.com/aadgraha/school_management/model/web"
	"github.com/aadgraha/school_management/service"
	"github.com/gin-gonic/gin"
)

type SubjectControllerImpl struct {
	subjectService service.SubjectService
}

func NewSubjectController(SubjectService service.SubjectService) SubjectController {
	return &SubjectControllerImpl{
		subjectService: SubjectService,
	}
}

func (controller *SubjectControllerImpl) FindById(c *gin.Context) {
	paramID := c.Param("subjectId")
	subjectResponses, err := controller.subjectService.FindById(paramID)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "find by id success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *SubjectControllerImpl) Create(c *gin.Context) {
	data_request := subject.InsertSubjectParams{}
	fmt.Println("c, &request", c, &data_request)
	helper.ReadFromRequestBody(c, &data_request)
	subjectResponses, err := controller.subjectService.Create(data_request)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "Create New Subject Success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *SubjectControllerImpl) Update(c *gin.Context) {
	paramID := c.Param("subjectId")
	data_request := subject.UpdateSubjectNewParams{}
	fmt.Println("c, &request", c, &data_request)
	helper.ReadFromRequestBody(c, &data_request)
	subjectResponses, err := controller.subjectService.Update(paramID, data_request)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "Edit Subject success",
		Data:    subjectResponses,
	}

	c.JSON(http.StatusOK, webResponse)

}

func (controller *SubjectControllerImpl) Delete(c *gin.Context) {
	paramID := c.Param("subjectId")
	subjectResponses, err := controller.subjectService.Delete(paramID)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: fmt.Sprintf("Delete Subject %s  Success", subjectResponses.Name),
		Data:    nil,
	}

	c.JSON(http.StatusOK, webResponse)
}
