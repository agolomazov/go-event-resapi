package models

import (
	"example.com/rest-api/db"
)

type Registration struct {
	ID      int64 `binding:"required" json:"id"`
	UserId  int64 `binding:"required" json:"user_id"`
	EventId int64 `binding:"required" json:"event_id"`
}

func (r *Registration) Save() (error) {
	query := `INSERT INTO 
	registrations (
		event_id,
		user_id
	) 
	VALUES (?, ?)`

	smtp, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer smtp.Close()

	result, err := smtp.Exec(r.EventId, r.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	
	r.ID = id
	return nil
}