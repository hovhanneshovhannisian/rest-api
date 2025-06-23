package dbnew

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBinit() {
	// Replace 'yourpassword' with your actual MySQL root password
	dsn := "root:Hovhannes20h@tcp(127.0.0.1:3306)/test"
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	// Verify the connection is valid
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Success!")

}
