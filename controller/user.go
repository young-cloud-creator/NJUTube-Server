package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nju-tube/security"
	"nju-tube/service"
	"nju-tube/structs"
	"strconv"
)

type userResponse struct {
	Response structs.Response `json:"response"`
	User     structs.User `json:"user"`
}

type userLRResponse struct {
	structs.Response `json:"response"`
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// UserInfo /user/ api handler
func UserInfo(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, userResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid User Id",
			},
		})
		return
	}

	tokenString := ctx.Query("token")
	if valid, uid := security.ValidateToken(tokenString); !valid || userId != uid {
		ctx.JSON(http.StatusOK, userResponse{
			Response: structs.Response{
				StatusCode: 2,
				StatusMsg:  "Invalid Token",
			},
		})
		return
	}

	user, err := service.QueryUserInfo(userId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, userResponse{
			Response: structs.Response{
				StatusCode: 3,
				StatusMsg:  "Unknown Error",
			},
		})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusOK, userResponse{
			Response: structs.Response{
				StatusCode: 4,
				StatusMsg:  "User Not Find",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, userResponse{
		Response: structs.Response{
			StatusCode: 0,
		},
		User: *user,
	})
}

// UserRegister /user/register/ api handler
func UserRegister(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	isSuccess, userId, msg := service.UserRegister(username, password)
	if !isSuccess {
		ctx.JSON(http.StatusOK, userLRResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  msg,
			},
		})
		return
	}

	tokenString, err := security.GenToken(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, userLRResponse{
			Response: structs.Response{
				StatusCode: 2,
				StatusMsg:  "Fail to Generate Token",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, userLRResponse{
		Response: structs.Response{
			StatusCode: 0,
		},
		UserId: userId,
		Token:  tokenString,
	})
}

// UserLogin /user/login/ api handler
func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	isSuccess, userId, msg := service.UserLogin(username, password)
	if !isSuccess {
		ctx.JSON(http.StatusOK, userLRResponse{
			Response: structs.Response{
				StatusCode: 1,
				StatusMsg:  msg,
			},
		})
		return
	}

	tokenString, err := security.GenToken(userId)
	if err != nil {
		ctx.JSON(http.StatusOK, userLRResponse{
			Response: structs.Response{
				StatusCode: 2,
				StatusMsg:  "Fail to Generate Token",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, userLRResponse{
		Response: structs.Response{
			StatusCode: 0,
		},
		UserId: userId,
		Token:  tokenString,
	})
}
