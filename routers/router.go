package routers

import (
	"example.com/rest-api/controllers"
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func Routers(server *gin.Engine) {
	server.GET("/posts", controllers.GetPosts)
	server.GET("/posts/:id", controllers.GetPostByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)
	authenticated.POST("/posts", controllers.CreatePost)
	authenticated.PUT("/posts/:id", controllers.UpdatePost)
	authenticated.DELETE("/posts/:id", controllers.DeletePost)

	server.POST("/register", controllers.Signup)
	server.POST("/login", controllers.Login)
}
