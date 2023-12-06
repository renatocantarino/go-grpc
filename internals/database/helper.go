package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "../../data.db")
	if err != nil {
		panic(err)
	}

	return db, nil

}
