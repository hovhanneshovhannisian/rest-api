package controllers

import (
	"example/blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(ctx *gin.Context) {
	var r_user models.User
	r_user.Username = ctx.Param("username")
	if isExist := r_user.IsExist(); !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "there is not such user",
		})
		return
	}
	var newMessage models.Message
	if err := ctx.BindJSON(&newMessage); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "missing parameters",
		})
		return
	}
	userID := ctx.GetInt64("userID")
	newMessage.SenderID = userID
	newMessage.ReceiverID = r_user.ID
	if err := newMessage.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error during sending",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message send!",
	})
}

func UpdateMessage(ctx *gin.Context) {

}

func GetMessages(ctx *gin.Context) {
	var s_user models.User
	s_user.Username = ctx.Param("username")
	if isExist := s_user.IsExist(); !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "there is not such user",
		})
		return
	}

	r_userID := ctx.GetInt64("userID")
	messages, err := models.GetChatMessages(r_userID, s_user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error during  chat fetching",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}
