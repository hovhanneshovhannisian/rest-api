package models

import (
	"example/blog/db"
	"time"
)

type Post struct {
	ID        int64
	Title     string `binding:"required"`
	Content   string `binding:"required"`
	AuthorID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) Save() error {
	query := `
	INSERT INTO posts
		(title, content, author_id)
	VALUES 
		(?, ?, ?)`
	result, err := db.DB.Exec(query, p.Title, p.Content, p.AuthorID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

func (p Post) Updated() error {
	query := `
	UPDATE posts
	SET title = ?, content = ?
	WHERE (id = ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Title, p.Content, p.ID)
	return err

}

func (p Post) Delete() error {
	query := "DELETE FROM posts WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID)
	return err
}

func GetPosts() ([]Post, error) {
	query := "SELECT * FROM posts"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostByID(id int64) (*Post, error) {
	query := `SELECT * FROM posts WHERE (id = ?)`
	result := db.DB.QueryRow(query, id)
	var post Post
	err := result.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
