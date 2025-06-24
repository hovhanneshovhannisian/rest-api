package main

import (
	"example/blog/db"
	"example/blog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routers.Router(server)

	server.Run(":8080")
}
