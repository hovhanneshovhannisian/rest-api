package controllers

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func GetPostByID(context *gin.Context) {
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	selected, err := models.GetPost(postID)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data": selected,
	})
}

func CreatePost(context *gin.Context) {
	var newPost models.Post
	if err := context.BindJSON(&newPost); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	authorID := context.GetInt64("authorID")
	newPost.Author = authorID
	if err := newPost.Save(); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "post has been created",
	})
}

func UpdatePost(context *gin.Context) {
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	authorID := context.GetInt64("authorID")
	post, err := models.GetPost(postID)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	if post.Author != authorID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "user is unauthorized",
		})
		return
	}
	var updatePost models.Post
	if err := context.BindJSON(&updatePost); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	updatePost.ID = postID
	updatePost.Author = authorID

	if err = updatePost.Update(); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "post has been updated",
		"data":    updatePost,
	})
}

func DeletePost(context *gin.Context) {
	postID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	authorID := context.GetInt64("authorID")
	post, err := models.GetPost(postID)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}

	if post.Author != authorID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "user is unauthorized",
		})
		return
	}

	if err := post.Delete(); err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "post has been deleted",
	})
}
