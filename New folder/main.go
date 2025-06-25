package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routers.Routers(server)
	//testing git user
	server.Run(":8080")
}
