package main

import (
	"example/blog/db"
	"example/blog/routers"

	_ "example/blog/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Blogging app
// @version         1.0
// @description     This is a sample server for a blogging platform. User signup/login and simple chatting feature with RESTapi

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	gotenv.Load()
	db.InitDB()

	server := gin.Default()
	server.Use(cors.Default())

	routers.Router(server)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8080")
}
