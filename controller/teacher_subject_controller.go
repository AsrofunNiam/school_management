package controller

import (
	"github.com/gin-gonic/gin"
)

type TeacherSubjectController interface {
	FindAll(context *gin.Context)
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
