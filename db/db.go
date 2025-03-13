package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to Database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
