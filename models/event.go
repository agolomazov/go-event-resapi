package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64 `json:"id"`
	Name        string `binding:"required" json:"name"`
	Description string `binding:"required" json:"description"`
	Location    string `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"dateTime"`
	UserId      int64 `json:"userId"`
}

func (e *Event) Save() (*Event, error) {
	query := `INSERT INTO 
	events (
		name,
		description,
		location,
		dateTime,
		user_id
	) 
	VALUES (?, ?, ?, ?, ?)`
	smtp, err := db.Db.Prepare(query)

	if err != nil {
		return &Event{}, err
	}

	defer smtp.Close()

	result, err := smtp.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return &Event{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return &Event{}, err
	}

	e.ID = id
	return e, nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.Db.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return &event, err
	}

	return &event, err
}

func (e *Event) Update() (error) {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`

	smtp, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e *Event) Delete() error {
	query := `
		DELETE FROM events WHERE id = ?
	`

	smtp, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	_, err = smtp.Exec(e.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) Unregister(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e *Event) ExistRegistration(userId int64) bool {
	query := "SELECT id, event_id, user_id FROM registrations WHERE event_id = ? AND user_id = ?"
	row := db.Db.QueryRow(query, e.ID, userId)

	var reg Registration

	err := row.Scan(&reg.ID, &reg.EventId, &reg.ID)

	if err != nil {
		return false
	}

	if reg.ID > 0 {
		return true
	}

	return false
}