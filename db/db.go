package db

import (
	"database/sql"

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
	createRegistrationsTable()
}

func createUsersTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	createTable(query, "Could not create users table.")
}

func createEventsTable() {
	query := `
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

	createTable(query, "Could not create events table.")
}

func createRegistrationsTable() {
	query := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT_NULL,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	createTable(query, "Could not create registration table.")
}

func createTable(query, errMsg string) {
	_, err := Db.Exec(query)
	checkQueryError(err, errMsg)
}