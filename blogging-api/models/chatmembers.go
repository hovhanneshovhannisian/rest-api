package models

import (
	"example/blog/db"
	"time"
)

type ChatMember struct {
	ID       int64
	RoomID   int64
	UserID   int64
	JoinedAt time.Time
}

func (m *ChatMember) Save() error {
	query := `
	INSERT INTO 
		(room_id, user_id)
	VALUES 
		(?, ?)`

	result, err := db.DB.Exec(query, m.RoomID, m.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = id
	return nil
}
