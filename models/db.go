package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	err = os.MkdirAll("db", os.ModePerm)

	if err != nil {
		log.Fatalf("Failed to create db directory: %v", err)
	}

	DB, err = sql.Open("sqlite3", "db/api.db")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Could not create users table: %v", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		userID INTEGER,
		FOREIGN KEY(userID) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id), 
		FOREIGN KEY(user_id) REFERENCES users(id) 
	)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		log.Fatalf("Could not create registrations table: %v", err)
	}
}
