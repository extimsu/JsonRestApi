package models

import (
	"errors"
	"time"

	"github.com/extimsu/JsonRestApi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {

	query := `
	INSERT INTO events (name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?);`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("error preparing statement " + err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return errors.New("error executing statement " + err.Error())
	}

	e.ID, err = result.LastInsertId()
	if err != nil {
		return errors.New("error with result ID " + err.Error())
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	var (
		events []Event
		event  Event
	)
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("cannot query events " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
			return nil, errors.New("cannot Scan event " + err.Error())
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID); err != nil {
		return nil, errors.New("cannot get event " + err.Error())
	}
	return &event, nil
}

func (event *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("Cannot prepare query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		return errors.New("Could not execute" + err.Error())
	}

	return nil
}

func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("cannot prepare query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	if err != nil {
		return errors.New("could not execute" + err.Error())
	}

	return nil
}

func (e *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("cannot prepare query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (e *Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("cannot prepare query: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
