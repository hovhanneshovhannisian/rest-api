package controllers

import (
	"example/blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var newPost models.Post
	if err := ctx.BindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	authorID := ctx.GetInt64("authorID")
	newPost.AuthorID = authorID
	if err := newPost.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "posted!",
	})
}

func GetAllPosts(ctx *gin.Context) {
	posts, err := models.GetPosts()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "no data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}

func GetPost(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	post, err := models.GetPostByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func UpdatePost(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	authorID := ctx.GetInt64("authorID")
	post, err := models.GetPostByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	if post.AuthorID != authorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}

	var updatingPost models.Post
	if err := ctx.BindJSON(&updatingPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	updatingPost.ID = post.ID

	if err := updatingPost.Updated(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "post updated!",
		"data":    updatingPost,
	})
}

func DeletePost(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	authorID := ctx.GetInt64("authorID")
	post, err := models.GetPostByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	if post.AuthorID != authorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}
	err = post.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "post deleted!",
	})
}
