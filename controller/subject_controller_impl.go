package controller

import (
	"net/http"

	sqlc_generate "github.com/aadgraha/school_management/model/sqlc"
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
	data_request := sqlc_generate.InsertSubjectParams{}
	controller.subjectService.Create(data_request)
}

func (controller *SubjectControllerImpl) Delete(c *gin.Context) {
	userId := 1
	controller.subjectService.Delete(int64(userId))
}

// func (controller *UserControllerImpl) FindAll(c *gin.Context, _ *auth.AccessDetails) {
// 	filters := helper.FilterFromQueryString(c, "dept.eq")

// 	userResponses := controller.userService.FindAll(&filters)
// 	webResponse := web.WebResponse{
// 		Success: true,
// 		Message: helper.MessageDataFoundOrNot(userResponses),
// 		Data:    userResponses,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }

func (controller *SubjectControllerImpl) FindById(c *gin.Context) {
	// request := "22"
	// filters := helper.FilterFromQueryString(c, "dept.eq")
	paramID := c.Param("customerId")

	// userResponse := controller.subjectService.FindById(request)
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
	data_request := sqlc_generate.UpdateSubjectNewParams{}

	controller.subjectService.Update(data_request)

}
