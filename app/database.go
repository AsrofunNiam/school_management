package app

import (
	"database/sql"

	db_query "github.com/aadgraha/school_management/model/sqlc"

	_ "github.com/lib/pq"
)

var TestQueries *db_query.Queries

func ConnectDatabase(user, host, password, port, db string) (*sql.DB, error) {

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	TestQueries = db_query.New(conn)

	return conn, nil
}
