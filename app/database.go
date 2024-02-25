package app

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// var TestQueries *db_query.Queries

func ConnectDatabase(user, host, password, port, db string) (*sql.DB, error) {

	dsn := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + db + "?sslmode=disable"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
