package routers

import (
	"example/blog/controllers"
	"example/blog/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/posts", controllers.GetAllPosts)
	server.GET("/posts/:id", controllers.GetPost)

	authentication := server.Group("/posts")
	authentication.Use(middlewares.Authentication)
	authentication.POST("/", controllers.CreatePost)
	authentication.POST("/:id/comment", controllers.CreateComment)
	authentication.PUT("/:id", controllers.UpdatePost)
	authentication.DELETE("/:id", controllers.DeletePost)

	server.GET("/comments", controllers.ToTestComments)

	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
