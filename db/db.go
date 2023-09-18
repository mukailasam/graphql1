package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnection() *sql.DB {
	dbConn, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic(err)
	}

	return dbConn
}
