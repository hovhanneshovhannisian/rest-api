package models

import (
	"example/blog/db"
	"time"
)

type Comment struct {
	ID        int64
	PostID    int64
	AuthorID  int64
	Content   string `binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
func (c Comment) Update() error {
	query := `
	UPDATE comments
	SET content = ?
	WHERE (id = ?)`
	_, err := db.DB.Exec(query, c.Content, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c Comment) Delete() error {
	query := "DELETE FROM comments WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.ID)
	return err
}

func GetCommentsByPostID(postID int64) ([]Comment, error) {
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
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func GetCommentByID(id int64) (*Comment, error) {
	query := `SELECT * FROM comments WHERE (id = ?)`
	result := db.DB.QueryRow(query, id)
	var comment Comment
	err := result.Scan(&comment.ID, &comment.PostID, &comment.AuthorID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
