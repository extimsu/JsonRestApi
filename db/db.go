package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Database connection pool.

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(5)
	if err := ceateTables(); err != nil {
		panic(err)
	}
}

func ceateTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		return errors.New("cannot create users table" + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		return errors.New("cannot create events table" + err.Error())
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		return errors.New("cannot create registrations table" + err.Error())
	}

	return nil
}
