package route

import (
	"database/sql"

	"github.com/aadgraha/school_management/controller"
	// dbx "github.com/aadgraha/school_management/model/sqlc"
	dbx "github.com/aadgraha/school_management/model"
	"github.com/aadgraha/school_management/repository"
	"github.com/aadgraha/school_management/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SubjectRoute(router *gin.Engine, dbx *dbx.Connect, db *sql.DB, validate *validator.Validate) {

	userService := service.NewSubjectServiceImpl(
		repository.NewSubjectRepository(dbx),
		db,
		// dbx,
		validate,
	)
	userController := controller.NewSubjectController(userService)

	router.GET("/subject/:customerId", userController.FindById)
	// router.GET("/call/:id", auth.Auth(dataController.FindByID, []string{}))
	// router.POST("/users", (userController.Create))
	// router.PUT("/users/:customerId", userController.Update, []string{})
	// router.GET("/users", userController.FindById, []string{})
	// router.DELETE("/users/:customerId", userController.Delete, []string{})
}
