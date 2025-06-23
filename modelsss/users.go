package modelsss

import (
	"example.com/rest-api/db"
	"example.com/rest-api/helper"
)

type User struct {
	ID       int64
	Username string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(username, email, password)
	VALUES (?, ?, ?)`

	hashedPass, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := db.DB.Exec(query, u.Username, u.Email, hashedPass)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	if err != nil {
		return err
	}
	return nil
}
