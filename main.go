package main

import (
	"fmt"
	"log"
	"net/http"

	app "github.com/aadgraha/school_management/app"
	c "github.com/aadgraha/school_management/configuration"
	"github.com/aadgraha/school_management/helper"
	conn "github.com/aadgraha/school_management/model"

	db_query "github.com/aadgraha/school_management/model/sqlc"
	"github.com/go-playground/validator/v10"
)

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

	// instance Connect
	conn := conn.Connect{
		Query: db_query.New(db),
	}

	router := app.NewRouter(db, &conn, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
