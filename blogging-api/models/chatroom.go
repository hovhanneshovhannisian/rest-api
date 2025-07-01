package models

import (
	"example/blog/db"
	"time"
)

type Chatroom struct {
	ID        int64
	Name      string `binding:"required"`
	Type      string
	CreatedAt time.Time
}

func (chat *Chatroom) Save() error {
	query := `
	INSERT INTO chatrooms
		(name, type)
	VALUES
		(?, ?)`

	// in further should add chat type either direct or group
	// in this stage its by default direct
	result, err := db.DB.Exec(query, chat.Name, "direct")
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	chat.ID = id
	return nil
}
