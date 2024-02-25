package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	// "github.com/AsrofunNiam/tch_indonesia_backend/app"
	app "github.com/aadgraha/school_management/app"
	c "github.com/aadgraha/school_management/configuration"
	"github.com/aadgraha/school_management/helper"

	db_query "github.com/aadgraha/school_management/model/sqlc"
	"github.com/go-playground/validator/v10"
)

var TestQueries *db_query.Queries

// var TestDB *sql.DB

func main() {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db, err := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)

	fmt.Println(err)

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	TestQueries = db_query.New(db)

	// routerTest := route_test.NewRouterSecond(db, validate)
	// serverSecond := http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: routerTest,
	// }

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	// err = serverSecond.ListenAndServe()
	helper.PanicIfError(err)
}
