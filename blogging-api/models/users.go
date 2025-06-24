package models

import (
	"errors"
	"example/blog/helper"

	"example/blog/db"
)

type User struct {
	ID       int64
	Username string `binding:"required"`
	Email    string
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

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users
	WHERE (username = ?)`
	row := db.DB.QueryRow(query, u.Username)
	var hashedPassword string
	if err := row.Scan(&u.ID, &hashedPassword); err != nil {
		return err
	}
	if ok := helper.VerifyPassword(u.Password, hashedPassword); !ok {
		return errors.New("invalid credentials")
	}
	return nil
}
