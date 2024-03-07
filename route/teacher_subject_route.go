package route

import (
	"database/sql"

	"github.com/aadgraha/school_management/controller"
	dbx "github.com/aadgraha/school_management/model"
	"github.com/aadgraha/school_management/repository"
	"github.com/aadgraha/school_management/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TeacherSubjectRoute(router *gin.Engine, dbx *dbx.Connect, db *sql.DB, validate *validator.Validate) {

	userService := service.NewTeacherSubjectServiceImpl(
		repository.NewTeacherSubjectRepository(),
		db,
		dbx,
		validate,
	)
	userController := controller.NewTeacherSubjectController(userService)

	router.GET("/teacher_subjects", userController.FindAll)
	router.GET("/teacher_subject/:teacherSubjectId", userController.FindById)
	router.POST("/teacher_subject", (userController.Create))
	router.PUT("/teacher_subject/:teacherSubjectId", userController.Update)
	router.DELETE("/teacher_subject/:teacherSubjectId", userController.Delete)
}
