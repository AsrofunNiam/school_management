package util

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	c "github.com/aadgraha/school_management/configuration"
	db "github.com/aadgraha/school_management/model/sqlc"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestObj(t *testing.T) {
	configuration, err := c.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	fmt.Println(configuration.AccessSecret + configuration.Db)

	testQueries = db.New(testDB)
}
