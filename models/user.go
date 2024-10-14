package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u User) Save() error {
	query := `
		INSERT INTO users(email, password) VALUES (?, ?)
	`
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.ID = userID
	return err
}