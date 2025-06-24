package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/helper"
)

type User struct {
	ID       int64
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(username, password)
	VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPass, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Username, hashedPass)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id //no need but assign the id that stores in the DB
	return err
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
