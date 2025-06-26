package models

import (
	"example/blog/db"
)

type Comment struct {
	ID       int64
	PostID   int64
	AuthorID int64
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

func GetComments(postID int64) ([]Comment, error) {
	qurey := `SELECT * FROM comments
	 WHERE (post_id = ?)`
	rows, err := db.DB.Query(qurey, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
}
