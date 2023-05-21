package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nju-tube/security"
	"nju-tube/service"
	"nju-tube/structs"
	"strconv"
)

type commentActionResponse struct {
	Response structs.Response `json:"response"`
	Comment  structs.Comment `json:"comment,omitempty"`
}

type commentListResponse struct {
	Response    structs.Response `json:"response"`
	CommentList []structs.Comment `json:"comment_list,omitempty"`
}

// CommentAction /comment/action/ api handler
func CommentAction(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.PostForm("video_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commentActionResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid Video Id",
			},
		})
		return
	}

	actionType, err := strconv.ParseInt(ctx.PostForm("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2) {
		ctx.JSON(http.StatusOK, commentActionResponse{
			Response: structs.Response{
				StatusCode: 2,
				StatusMsg:  "Invalid Action Type",
			},
		})
		return
	}

	tokenString := ctx.PostForm("token")
	valid, userId := security.ValidateToken(tokenString)
	if !valid {
		ctx.JSON(http.StatusOK, commentActionResponse{
			Response: structs.Response{
				StatusCode: 3,
				StatusMsg:  "Invalid Token",
			},
		})
		return
	}

	switch actionType {
	case 1:
		commentText := ctx.PostForm("comment_text")
		comment, err := service.AddComment(userId, videoId, commentText)
		if err != nil {
			ctx.JSON(http.StatusOK, commentActionResponse{
				Response: structs.Response{
					StatusCode: 4,
					StatusMsg:  "Database Error",
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, commentActionResponse{
			Response: structs.Response{
				StatusCode: 0,
			},
			Comment: *comment,
		})

	case 2:
		commentId, err := strconv.ParseInt(ctx.PostForm("comment_id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, commentActionResponse{
				Response: structs.Response{
					StatusCode: 5,
					StatusMsg:  "Invalid Comment Id",
				},
			})
			return
		}

		success, msg := service.DeleteComment(userId, videoId, commentId)
		if success {
			ctx.JSON(http.StatusOK, commentActionResponse{
				Response: structs.Response{
					StatusCode: 0,
				}})
		} else {
			ctx.JSON(http.StatusOK, commentActionResponse{
				Response: structs.Response{
					StatusCode: 6,
					StatusMsg:  msg,
				}})
		}
	}
}

// CommentList /comment/list/ api handler
func CommentList(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, commentListResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid Video Id",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, commentListResponse{
		Response:    structs.Response{StatusCode: 0},
		CommentList: service.CommentList(videoId),
	})
}
