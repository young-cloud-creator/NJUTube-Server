package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nju-tube/security"
	"nju-tube/service"
	"nju-tube/structs"
	"strconv"
)

// FavoriteAction /favorite/action/ api handler
func FavoriteAction(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.PostForm("video_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid Video Id",
		})
		return
	}

	actionType, err := strconv.ParseInt(ctx.PostForm("action_type"), 10, 64)
	if err != nil || (actionType != 1 && actionType != 2 && actionType != 3) {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 2,
			StatusMsg:  "Invalid Action Type",
		})
		return
	}

	tokenString := ctx.PostForm("token")
	valid, userId := security.ValidateToken(tokenString)
	if !valid {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 3,
			StatusMsg:  "Invalid Token",
		})
		return
	}

	if actionType == 3 {
		isFavorite, err := service.UserFavoriteVideo(userId, videoId)
		if err != nil {
			ctx.JSON(http.StatusOK, structs.Response{
				StatusCode: 5,
				StatusMsg:  "Fail to Get Info",
			})
			return
		}
		var status int32 = 0
		if isFavorite {
			status = 1
		}
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: status,
		})
		return
	}

	err = service.DoFavoriteAction(userId, videoId, actionType == 1)
	if err != nil {
		ctx.JSON(http.StatusOK, structs.Response{
			StatusCode: 4,
			StatusMsg:  "Fail to Action",
		})
		return
	}

	ctx.JSON(http.StatusOK, structs.Response{StatusCode: 0})
}

// FavoriteList /favorite/list/ api handler
func FavoriteList(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, videoListResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid User Id",
			},
		})
		return
	}

	tokenString := ctx.Query("token")
	valid, selfId := security.ValidateToken(tokenString)
	if !valid {
		selfId = 0
	}

	videoList, err := service.FavoriteList(userId, selfId)
	if err != nil {
		ctx.JSON(http.StatusOK, videoListResponse{
			Response: structs.Response{
				StatusCode: 3,
				StatusMsg:  "Error Occurred in Database",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, videoListResponse{
		Response:  structs.Response{StatusCode: 0},
		VideoList: videoList,
	})
}
