package controllers

import (
	"example/blog/helper"
	"example/blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Signup godoc
// @Summary      Sign-up user
// @Description  Create a user
// @Tags         auth
// @Param        user  body models.User  true  "User info"
// @Accept json
// @Router       /signup [post]
func SignUp(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	if isExist := newUser.IsExist(); isExist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "user already exists",
		})
		return
	}
	if err := newUser.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user has been successfully created",
	})
}

// Login godoc
// @Summary      Login user
// @Description  User login
// @Tags         auth
// @Param        user  body models.User  true  "User Credentials"
// @Accept json
// @Router       /login [post]
func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	if err := user.ValidateCredentials(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	access_T, refresh_T, err := helper.GenerateToken(user.Username, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.SetCookie("refreshtoken", refresh_T,
		7200,
		"/",
		"http://localhost:3000",
		true,
		true)
	ctx.JSON(http.StatusOK, gin.H{
		"message":     "Logged in!",
		"accesstoken": access_T,
	})
}
