package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "blog.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTabel()
}

// func InitDB() {
// 	cfg := mysql.NewConfig()
// 	cfg.User = os.Getenv("DBUSER")
// 	cfg.Passwd = os.Getenv("DBPASS")
// 	cfg.Net = "tcp"
// 	cfg.Addr = "127.0.0.1:3306"
// 	cfg.DBName = "recordings"

// 	// Get a database handle.
// 	var err error
// 	DB, err = sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := DB.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected!")

// 	DB.SetMaxOpenConns(10)
// 	DB.SetMaxIdleConns(5)

// 	createTabel()
// }

func createTabel() {
	// Users table
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT,
		password TEXT NOT NULL
	);`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("couldn't create users table")
	}

	// Posts table
	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author_id) REFERENCES users(id)
	);`

	//	TODO
	//	after implementing the mysql not sqlite change the
	// 	table creation query to this to work update_at time
	// 	setting automaticlly sqlite dot not support this query
	//
	//
	// `CREATE TABLE IF NOT EXISTS posts (
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	title TEXT NOT NULL,
	// 	content TEXT NOT NULL,
	// 	author_id INTEGER NOT NULL,
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 	// updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, //
	// 	FOREIGN KEY (author_id) REFERENCES users(id)
	// );`

	_, err = DB.Exec(createPostsTable)
	if err != nil {
		panic("couldn't create posts table")
	}

	// Comments table
	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		post_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id),
		FOREIGN KEY (author_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createCommentsTable)
	if err != nil {
		panic("couldn't create comments table")
	}
}
