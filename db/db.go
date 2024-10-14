package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable()
	createEventsTable()
}

func createUsersTable() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := Db.Exec(createUsersTable)

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create users table.")
	}
}

func createEventsTable() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err := Db.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not create events table.")
	}
}