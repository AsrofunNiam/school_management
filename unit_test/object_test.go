package util

import (
	"database/sql"
	"fmt"

	// "fmt"
	"log"

	// "fmt"
	"testing"
	// "github.com/aadgraha/school_management/exception"
	// app "github.com/aadgraha/school_management/app"
	c "github.com/aadgraha/school_management/configuration"
	db "github.com/aadgraha/school_management/model/sqlc"
	// "github.com/gin-gonic/gin"
	// "github.com/go-playground/validator/v10"
)

// func ErrorHandlerTest() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		defer func() {
// 			err := recover()
// 			fmt.Println(err)
// 			if err != nil {
// 				// fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
// 				exception.ErrorHandler(c, err)
// 			}
// 		}()
// 		c.Next()
// 	}
// }

// func NewRouterSecond(db *sql.DB, validate *validator.Validate) *gin.Engine {
// 	// routing data
// 	router := gin.New()
// 	router.Use(ErrorHandlerTest())

// 	return router
// }

var testQueries *db.Queries
var testDB *sql.DB

// func ConnectDatabase(user, host, password, port, db string) *sql.DB {

// 	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
// 	conn, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return nil
// 	}

// 	return conn
// }

func TestObj(t *testing.T) {
	configuration, err := c.LoadConfig()
	// ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println(configuration.AccessSecret + configuration.Db)

	testQueries = db.New(testDB)
}
