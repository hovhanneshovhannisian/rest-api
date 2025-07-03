package main

import (
	"example/blog/db"
	"example/blog/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	db.InitDB()
	server := gin.Default()
	server.Use(cors.Default())

	routers.Router(server)

	server.Run(":8080")
}
