package app

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase(user, host, password, port, db string) (*sql.DB, error) {

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
