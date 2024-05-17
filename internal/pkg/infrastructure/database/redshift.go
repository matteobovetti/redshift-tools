package database

import (
	"database/sql"
	"fmt"
	"log"
)

func connectionStringBuilder(username, password, hostname, port, database string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, hostname, port, database)
}

func CreateConnection(username, password, hostname, port, database string) *sql.DB {
	db, err := sql.Open("postgres", connectionStringBuilder(username, password, hostname, port, database))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
