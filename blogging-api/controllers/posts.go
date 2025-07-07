package controllers

import (
	"example/blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePost godoc
// @Summary      Create a post
// @Description  Create a new blog post
// @Tags         posts
// @Param        title  body string  true  "Post Title"
// @Param        content  body string  true  "Post Content"
// @Param 		 Authorization header string 	true "Access token"
// @Success      200   {object}  models.Post
// @Router       /posts [post]
func CreatePost(ctx *gin.Context) {
	var newPost models.Post
	if err := ctx.BindJSON(&newPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	authorID := ctx.GetInt64("userID")
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

// GetAllPosts godoc
// @Summary      Get all the posts
// @Description  Fetching all the post in the platform
// @Tags         posts
// @Success      200   {array}  models.Post
// @Router       /posts [get]
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

// GetPost godoc
// @Summary      Get the post
// @Description  Fetching the post by post ID
// @Tags         posts
// @Param 		 postId path string true "Post ID"
// @Success      200   {object}  models.Post
// @Router       /posts/{postId} [get]
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

// UpdatePost godoc
// @Summary      Update the post
// @Description  Updating the post
// @Tags         posts
// @Param 		 postId path string true "Post ID"
// @Param        title  body string  true  "Post Title"
// @Param        content  body string  true  "Post Content"
// @Param 		 Authorization header string 	true "Access token"
// @Success      200   {object}  models.Post
// @Router       /posts/{postId} [put]
func UpdatePost(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	authorID := ctx.GetInt64("userID")
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

// DeletePost godoc
// @Summary      Delete the post
// @Description  Deleting the post
// @Tags         posts
// @Param 		 postId path string true "Post ID"
// @Param 		 Authorization header string true "Access token"
// @Success      200   {object}  models.Post
// @Router       /posts/{postId} [delete]
func DeletePost(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	authorID := ctx.GetInt64("userID")
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
