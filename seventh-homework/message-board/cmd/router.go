package cmd

import (
	"github.com/gin-gonic/gin"
	"message-board/api"
)

func AccessUrl() {
	router := gin.Default()

	//user`s url
	router.POST("/register", api.Register)
	router.GET("/login", api.Login)
	router.PUT("/updatePwd", api.UpdatePwd)

	//message`s url
	router.POST("/addLeavingMessage", api.AddLeavingMessage)
	router.POST("/addCommentMessage", api.AddCommentMessage)
	router.GET("/selectAllLeavingMessages", api.SelectAllLeavingMessages)
	router.GET("/selectCommentMessages", api.SelectCommentMessages)
	router.PUT("/updateLeavingMessage", api.UpdateLeavingMessage)
	router.PUT("/updateCommentMessage", api.UpdateCommentMessage)
	router.DELETE("/deleteLeavingMessage", api.DeleteLeavingMessage)
	router.DELETE("/deleteCommentMessage", api.DeleteCommentMessage)

	//监听端口
	router.Run(":8081")
}
