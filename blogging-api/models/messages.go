package models

import (
	"example/blog/db"
	"log"
	"time"
)

type Message struct {
	ID         int64
	SenderID   int64
	ReceiverID int64
	Message    string `binding:"required"`
	IsRead     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m *Message) Save() error {
	query := `
	INSERT INTO messages
		(sender_id, receiver_id, message)
	VALUES
		(?, ?, ?)`

	result, err := db.DB.Exec(query, m.SenderID, m.ReceiverID, m.Message)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	m.ID = id
	return nil
}

func (m *Message) Update() error {
	query := `
	UPDATE comments
	SET	message = ?
	WHERE (id = ?)`

	result, err := db.DB.Exec(query, m.Message, m.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	m.ID = id
	return nil
}

func GetChatMessages(s, r int64) ([]Message, error) {
	query := `
	SELECT * FROM messages
	WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
	ORDER BY created_at`

	rows, err := db.DB.Query(query, s, r, r, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Message, &message.IsRead, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func GetMessageByID(id int64) (*Message, error) {
	query := "SELECT * FROM messages WHERE (id = ?)"

	row := db.DB.QueryRow(query, id)
	var message Message
	err := row.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.Message, &message.IsRead, &message.CreatedAt, &message.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
