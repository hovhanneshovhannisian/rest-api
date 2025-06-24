package middlewares

import (
	"net/http"

	"example.com/rest-api/helper"
	"github.com/gin-gonic/gin"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user is unauthorized"})
		return
	}
	authorID, err := helper.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	context.Set("authorID", authorID)
	context.Next()
}
