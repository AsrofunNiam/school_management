package main

// "github.com/AsrofunNiam/tch_indonesia_backend/app"
// "github.com/AsrofunNiam/tch_indonesia_backend/helper"

// func main() {
// 	configuration, err := c.LoadConfig()
// 	if err != nil {
// 		log.Fatalln("Failed at config", err)
// 	}

// 	port := configuration.Port
// 	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)

// 	skiDb := app.ConnectDatabase(configuration.UserSki, configuration.HostSki, configuration.PasswordSki, configuration.PortDBSki, configuration.DbSki)

// 	// Validator
// 	validate := validator.New()
// 	helper.RegisterValidation(validate)

// 	router := app.NewRouter(db, validate, skiDb)
// 	server := http.Server{
// 		Addr:    ":" + port,
// 		Handler: router,
// 	}
// 	log.Printf("Server is running on port %s", port)

// 	err = server.ListenAndServe()
// 	helper.PanicIfError(err)
// }
