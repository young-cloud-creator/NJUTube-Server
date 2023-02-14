package controller

import (
	"github.com/gin-gonic/gin"
	"goto2023/service"
	"goto2023/structs"
	"net/http"
	"strconv"
	"time"
)

type feedResponse struct {
	Response  structs.Response
	NextTime  int64           `json:"next_time"`
	VideoList []structs.Video `json:"video_list,omitempty"`
}

// FeedList /feed/ api handler
func FeedList(ctx *gin.Context) {
	latestTime, err := strconv.ParseInt(ctx.Query("latest_time"), 10, 64)
	if err != nil || latestTime == 0 {
		latestTime = time.Now().Unix() * 1000 // second to millisecond
	}
	latestTime = latestTime / 1000 // millisecond to second

	videoList, nextTime, err := service.GenerateFeed(latestTime)
	if err != nil {
		ctx.JSON(http.StatusOK, feedResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, feedResponse{
		Response: structs.Response{
			StatusCode: 0,
		},
		NextTime:  nextTime * 1000, // second to millisecond
		VideoList: videoList,
	})
}
