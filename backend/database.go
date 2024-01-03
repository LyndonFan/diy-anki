package backend

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "backend/database.db"
const createTableFile = "backend/createTables.sql"

func GetDataBase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	err = createTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	data, err := os.ReadFile(createTableFile)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}
	return nil
}
