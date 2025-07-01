package controllers

import (
	"example/blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChatRoom(ctx *gin.Context) {
	var newChat models.Chatroom
	if err := ctx.BindJSON(&newChat); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "missing paramets",
		})
		return
	}

	if err := newChat.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error during chat creating",
		})
		return
	}
	userID := ctx.GetInt64("userID")
	var newMember models.ChatMember
	newMember.RoomID = newChat.ID
	newMember.UserID = userID
	if err := newMember.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error during chat member saving",
		})
		return
	}
}
