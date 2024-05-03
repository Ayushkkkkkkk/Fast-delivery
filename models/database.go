package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB // Exported variable

func InitDB() error {
	db, err := sql.Open("sqlite3", "fast-delivery.db")
	if err != nil {
		return err
	}
	Db = db // Assigning to exported variable
	CreateUserTable()
	return nil
}
