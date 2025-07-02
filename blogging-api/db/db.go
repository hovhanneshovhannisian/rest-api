package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "blogging"
	cfg.ParseTime = true

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTabel()
}

func createTabel() {
	// Users table
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255),
		password VARCHAR(255) NOT NULL
	);`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("couldn't create users table")
	}

	// Posts table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		content VARCHAR(255) NOT NULL,
		author_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (author_id) REFERENCES users(id)
	);`
	_, err = DB.Exec(createPostsTable)
	if err != nil {
		panic("couldn't create posts table")
	}

	// Comments table
	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INT PRIMARY KEY AUTO_INCREMENT,
		post_id INT NOT NULL,
		author_id INT NOT NULL,
		content VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (author_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createCommentsTable)
	if err != nil {
		panic("couldn't create comments table")
	}

	// messages table
	createMessagesTable := `
	CREATE TABLE IF NOT EXISTS messages (
		id INT PRIMARY KEY AUTO_INCREMENT,
		sender_id INT NOT NULL,
		receiver_id INT NOT NULL,
		message VARCHAR(255) NOT NULL,
		is_read BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (sender_id) REFERENCES users(id),
		FOREIGN KEY (receiver_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createMessagesTable)
	if err != nil {
		log.Fatal(err)
		panic("couldn't create messages table")
	}

	// chatrooms table
	// createChatRoomsTable := `
	// CREATE TABLE IF NOT EXISTS chatrooms (
	// 	room_id INT PRIMARY KEY AUTO_INCREMENT,
	// 	name VARCHAR(45) NOT NULL,
	// 	type ENUM('direct', 'group') NOT NULL,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// )`
	// _, err = DB.Exec(createChatRoomsTable)
	// if err != nil {
	// 	panic("couldn't create chatrooms table")
	// }

	// chatmembers table
	// createChatMembersTable := `
	// CREATE TABLE IF NOT EXISTS chatmembers (
	// 	id INT PRIMARY KEY AUTO_INCREMENT,
	// 	room_id INT NOT NULL,
	// 	user_id INT NOT NULL,
	// 	joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	// 	FOREIGN KEY (room_id) REFERENCES chatmembers(room_id),
	// 	FOREIGN KEY (user_id) REFERENCES users(id)
	// )`
	// _, err = DB.Exec(createChatMembersTable)
	// if err != nil {
	// 	panic("couldn't create chatmembers table")
	// }
}
