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
	authentication.PUT("/:id", controllers.UpdatePost)
	authentication.DELETE("/:id", controllers.DeletePost)

	authentication.POST("/:id/comment", controllers.CreateComment)
	authentication.GET("/:id/comment", controllers.GetPostComments)

	authentication.PUT("/comment/:id", controllers.UpdateComment)
	authentication.DELETE("/comment/:id", controllers.DeleteComment)

	//server.GET("/comments", controllers.ToTestComments)

	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
