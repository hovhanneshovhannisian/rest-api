package modelsss

import "example.com/rest-api/db"

type Post struct {
	ID       int64
	Title    string `binding:"required"`
	Content  string `binding:"required"`
	AuthorID int64  `binding:"required"`
}

func (p Post) Save() error {
	query := `
	INSERT INTO posts
		(title, content, author_id)
	VALUES 
		(?, ?, ?)`
	result, err := db.DB.Exec(query, p.Title, p.Content, p.AuthorID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
