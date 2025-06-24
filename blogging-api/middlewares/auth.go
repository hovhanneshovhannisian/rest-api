package middlewares

import (
	"example/blog/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user is unauthorized"})
		return
	}
	authorID, err := helper.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
		return
	}

	ctx.Set("authorID", authorID)
	ctx.Next()
}
