package controllers

import (
	"example/blog/db"
	"example/blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PublishComment godoc
// @Summary      Write a comment
// @Description  Commenting the post
// @Tags         posts
// @Param 		 postId path string true "Post ID"
// @Param        comment  body models.Comment  true  "Comment content"
// @Param 		 Authorization header string 	true "Access token"
// @Accept 		 json
// @Produce 	 json
// @Router       /posts/{postID}/comment [post]
func CreateComment(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	_, err = models.GetPostByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	authorID := ctx.GetInt64("userID")

	var newComment models.Comment
	if err := ctx.BindJSON(&newComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	newComment.AuthorID = authorID
	newComment.PostID = cnvtID
	if err := newComment.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "commented!",
	})
}

// UpdateComment godoc
// @Summary      Update the comment
// @Description  Updating the comment
// @Tags         posts
// @Param 		 commentID path string true "Comment ID"
// @Param        comment  body models.Comment  true  "Comment content"
// @Param 		 Authorization header string 	true "Access token"
// @Accept 		 json
// @Produce 	 json
// @Router       /comment/{commentID} [put]
func UpdateComment(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	cmnt, err := models.GetCommentByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "comment not found",
		})
		return
	}
	authorID := ctx.GetInt64("userID")
	if cmnt.AuthorID != authorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}
	var updatedComment models.Comment
	if err := ctx.BindJSON(&updatedComment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing parameters",
		})
		return
	}
	updatedComment.ID = cmnt.ID
	if err := updatedComment.Update(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "comment updated!",
	})
}

// DeleteComment godoc
// @Summary      Delete the comment
// @Description  Deleting the comment
// @Tags         posts
// @Param 		 commentID path string true "Comment ID"
// @Param 		 Authorization header string 	true "Access token"
// @Produce 	 json
// @Router       /comment/{commentID} [delete]
func DeleteComment(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	cmnt, err := models.GetCommentByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "comment not found",
		})
		return
	}
	authorID := ctx.GetInt64("userID")
	if cmnt.AuthorID != authorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized user",
		})
		return
	}
	err = cmnt.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "comment deleted!",
	})
}

// GetPostComments godoc
// @Summary      Fetching the post comments
// @Description  Fetching all the comments of the post
// @Tags         posts
// @Param 		 postId path string true "Post ID"
// @Param 		 Authorization header string 	true "Access token"
// @Produce 	 json
// @Router       /posts/{postID}/comment [get]
func GetPostComments(ctx *gin.Context) {
	cnvtID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	_, err = models.GetPostByID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "post not found",
		})
		return
	}
	cmnts, err := models.GetCommentsByPostID(cnvtID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "comment not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": cmnts,
	})
}

// to be removed
func ToTestComments(ctx *gin.Context) {

	query := "SELECT id, post_id, author_id, content FROM comments"
	rows, err := db.DB.Query(query)
	if err != nil {
		ctx.JSON(408, gin.H{
			"error": err,
		})
		return
	}
	defer rows.Close()
	var posts []models.Comment
	for rows.Next() {
		var post models.Comment
		err := rows.Scan(&post.ID, &post.PostID, &post.AuthorID, &post.Content)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err,
			})
			return
		}
		posts = append(posts, post)
	}
	ctx.JSON(200, gin.H{
		"data": posts,
	})
}
