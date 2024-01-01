package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "database.db"

func getDataBase() (*sql.DB, error) {
	return sql.Open("sqlite3", file)
}
