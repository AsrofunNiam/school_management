package app

import (
	"database/sql"
	"fmt"

	"github.com/aadgraha/school_management/exception"
	"github.com/aadgraha/school_management/route"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	// db_query "github.com/aadgraha/school_management/model/sqlc"
)

// var TestQueries *db_query.Queries

// var TestDB *sql.DB

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			fmt.Println(err)
			if err != nil {
				// fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				exception.ErrorHandler(c, err)
			}
		}()
		c.Next()
	}
}

func NewRouter(db *sql.DB, validate *validator.Validate) *gin.Engine {
	// routing data
	// TestQueries = db_query.New(db)
	router := gin.New()
	router.Use(ErrorHandler())
	route.SubjectRoute(router, db, TestQueries, validate)
	return router
}
