package main

import (
	"github.com/gin-gonic/gin"
	"goto2023/controller"
	"net/http"
)

func initRouter(router *gin.Engine) {
	// resources like videos, covers, ...
	router.StaticFS("/resources", http.Dir("./public"))

	// apis like feed, login, ...
	apiRouter := router.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.FeedList)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.UserRegister)
	apiRouter.POST("/user/login/", controller.UserLogin)
	apiRouter.POST("/publish/action/", controller.PublishAction)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	/*
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
		apiRouter.GET("/relation/friend/list/", controller.FriendList)
		apiRouter.GET("/message/chat/", controller.MessageChat)
		apiRouter.POST("/message/action/", controller.MessageAction)
	*/
}
