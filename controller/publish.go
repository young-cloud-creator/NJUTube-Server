package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goto2023/security"
	"goto2023/service"
	"goto2023/structs"
	"net/http"
	"path/filepath"
	"strconv"
)

type publishListResponse struct {
	Response  structs.Response
	VideoList []structs.Video `json:"video_list"`
}

// PublishAction /publish/action/ api handler
func PublishAction(ctx *gin.Context) {
	tokenString := ctx.PostForm("token")
	valid, userId := security.ValidateToken(tokenString)
	if !valid {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid Token",
		})
		return
	}

	title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	saveName := fmt.Sprintf("%d-%s", userId, filename)
	fullName := filepath.Join(service.VideoDir, saveName)
	if err = ctx.SaveUploadedFile(data, fullName); err != nil {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 3,
			StatusMsg:  "cannot save video file",
		})
		return
	}

	if err = service.PublishAction(title, saveName, userId); err != nil {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 4,
			StatusMsg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, structs.Response{
		StatusCode: 0,
	})
}

// PublishList /publish/list/ api handler
func PublishList(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, publishListResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid User Id",
			},
		})
		return
	}

	tokenString := ctx.Query("token")
	if valid, uid := security.ValidateToken(tokenString); !valid || userId != uid {
		ctx.JSON(http.StatusOK, publishListResponse{
			Response: structs.Response{
				StatusCode: 2,
				StatusMsg:  "Invalid Token",
			},
		})
		return
	}

	videos, err := service.PublishList(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, publishListResponse{
			Response: structs.Response{
				StatusCode: 3,
				StatusMsg:  "Unknown Error",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, publishListResponse{
		Response: structs.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
