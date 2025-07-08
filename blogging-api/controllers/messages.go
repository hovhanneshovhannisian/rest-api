package controllers

import (
	"example/blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SendMessage godoc
// @Summary      Send Message
// @Description  Sending message to existing user
// @Tags         messages
// @Param        username  path string  true  "the reciever user"
// @Param        message  body models.Message  true  "Message content"
// @Param 		 Authorization header string 	true "Access token"
// @Accept 		 json
// @Produce 	 json
// @Router       /message/{username} [post]
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

// UpdateMessage godoc
// @Summary      Update Message
// @Description  Edit wrote message
// @Tags         messages
// @Param        messageID  path string  true  "Message ID"
// @Param        message  body models.Message  true  "Message content"
// @Param 		 Authorization header string 	true "Access token"
// @Accept 		 json
// @Produce 	 json
// @Router       /message/{messageID} [put]
func UpdateMessage(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	senderID := ctx.GetInt64("userID")
	message, err := models.GetMessageByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	if message.SenderID != senderID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "you can't edit this message",
		})
		return
	}

	var editmessage models.Message
	if err := ctx.BindJSON(&editmessage); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	editmessage.ID = cnvtID
	if err := editmessage.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "message edited!",
		"data":    editmessage,
	})
}

// GetAllMessage godoc
// @Summary      Get messages
// @Description  Getting the chat messages
// @Tags         messages
// @Param        username  path string  true  "the reciever user"
// @Param 		 Authorization header string 	true "Access token"
// @Produce 	 json
// @Router       /message/{username} [get]
func GetMessages(ctx *gin.Context) {
	var r_user models.User
	r_user.Username = ctx.Param("username")
	if isExist := r_user.IsExist(); !isExist {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "there is no such user",
		})
		return
	}

	s_userID := ctx.GetInt64("userID")
	messages, err := models.GetChatMessages(s_userID, r_user.ID)
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
