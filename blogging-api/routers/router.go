package routers

import (
	"example/blog/controllers"
	"example/blog/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {

	apiv1 := server.Group("/api/v1")
	//signup and login
	apiv1.POST("/signup", controllers.SignUp)
	apiv1.POST("/login", controllers.Login)
	{
		posts := apiv1.Group("/posts")
		{
			posts.GET("/", controllers.GetAllPosts)
			posts.GET("/:id", controllers.GetPost)

			//only authenticated
			posts.Use(middlewares.Authentication)

			posts.POST("/", controllers.CreatePost)
			posts.PUT("/:id", controllers.UpdatePost)
			posts.DELETE("/:id", controllers.DeletePost)

			posts.POST("/:id/comment", controllers.CreateComment)
			posts.GET("/:id/comment", controllers.GetPostComments)
			posts.PUT("/comment/:id", controllers.UpdateComment)
			posts.DELETE("/comment/:id", controllers.DeleteComment)
		}

		// live chat

		message := apiv1.Group("/message")
		message.Use(middlewares.Authentication)
		{
			message.POST("/:username", controllers.SendMessage)
			message.GET("/:username", controllers.GetMessages)
			message.PUT("/:id", controllers.UpdateMessage)
		}
	}
}
