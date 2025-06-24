package routers

import (
	"example/blog/controllers"
	"example/blog/middlewares"

	"github.com/gin-gonic/gin"
)

func Router(server *gin.Engine) {
	server.GET("/posts", controllers.GetAllPosts)
	server.GET("/posts/:id", controllers.GetPost)
	server.POST("/posts", middlewares.Authentication, controllers.CreatePost)
	server.PUT("/posts/:id", middlewares.Authentication, controllers.UpdatePost)

	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
