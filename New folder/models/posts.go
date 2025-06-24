package models

import (
	"example.com/rest-api/db"
)

type Post struct {
	ID          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Author      int64
}

func (p *Post) Save() error {
	query := `
	INSERT INTO posts(title, description, author)
	VALUES (?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Title, p.Description, p.Author)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.ID = id
	return err
}

func (p *Post) Update() error {
	query := `
	UPDATE posts
	SET title = ?, description = ?, author = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Title, p.Description, p.Author, p.ID)
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

func GetAllPosts() ([]Post, error) {
	query := `SELECT * FROM posts`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Author)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(id int64) (*Post, error) {
	query := `SELECT * FROM posts WHERE (id = ?)`
	row := db.DB.QueryRow(query, id)
	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Description, &post.Author)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
