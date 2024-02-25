package controller

import (
	"github.com/gin-gonic/gin"
)

type SubjectController interface {
	FindById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}
