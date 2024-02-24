package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/AsrofunNiam/tch_indonesia_backend/app"
	app "github.com/aadgraha/school_management/app"
	c "github.com/aadgraha/school_management/configuration"
	"github.com/aadgraha/school_management/helper"
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

	// 	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}