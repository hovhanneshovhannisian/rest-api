package modelsss

import "example.com/rest-api/db"

type Comment struct {
	ID       int64
	PostID   int64  `binding:"required"`
	AuthorID int64  `binding:"required"`
	Content  string `binding:"required"`
}

func (c Comment) Save() error {
	query := `
	INSERT INTO comments
		(post_id, author_id, content)
	VALUES 
		(?, ?, ?)`
	result, err := db.DB.Exec(query, c.PostID, c.AuthorID, c.Content)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}
