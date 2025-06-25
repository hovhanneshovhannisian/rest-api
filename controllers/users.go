package controllers

import (
	"net/http"

	"example.com/rest-api/helper"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := newUser.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "user has  created",
	})
}

func Login(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"esrror": err,
		})
		return
	}
	token, err := helper.GenerateToken(user.Username, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"errodr": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "you are logged in",
		"token":   token,
	})
}
