package controller

import (
	"net/http"

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

func (controller *SubjectControllerImpl) Create(c *gin.Context) {
	data_request := subject.InsertSubjectParams{}
	controller.subjectService.Create(data_request)
}

func (controller *SubjectControllerImpl) Delete(c *gin.Context) {
	userId := 1
	controller.subjectService.Delete(int64(userId))
}

func (controller *SubjectControllerImpl) FindById(c *gin.Context) {
	paramID := c.Param("subjectId")
	userResponses, err := controller.subjectService.FindById(paramID)
	if err != nil {
		return
	}

	webResponse := web.WebResponse{
		Success: true,
		Message: "find by id success",
		Data:    userResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *SubjectControllerImpl) Update(c *gin.Context) {
	data_request := subject.UpdateSubjectNewParams{}

	controller.subjectService.Update(data_request)

}
