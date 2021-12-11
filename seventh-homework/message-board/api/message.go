package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/models"
	"message-board/service"
)

func AddLeavingMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	//fmt.Println(m)
	err = service.AddLeavingMessage(m.UserName, m.Id, m.Content)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Add Leaving Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "add leaving message succeed!",
		})
	}
}

func AddCommentMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	err = service.AddCommentMessage(m.UserName, m.PId,m.CId, m.Content)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Add Comment Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "add Comment message succeed!",
		})
	}
}

func UpdateLeavingMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	err = service.UpdateLeavingMessage(m.UserName, m.Id, m.Content)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Update Leaving Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "Update leaving message succeed!",
		})
	}
}

func UpdateCommentMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	err = service.UpdateCommentMessage(m.UserName, m.CId, m.Content)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Update Comment Message Failed...",
			"error info": err,
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "Update Comment message succeed!",
		})
	}
}

func DeleteLeavingMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	err = service.DeleteLeavingMessage(m.UserName, m.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Delete Leaving Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "Delete Leaving message succeed!",
		})
	}
}

func DeleteCommentMessage(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	err = service.DeleteCommentMessage(m.UserName, m.CId)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Delete Leaving Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"MSG": "Delete Leaving message succeed!",
		})
	}
}

func SelectAllLeavingMessages(c *gin.Context) {
	messages, err := service.SelectAllLeavingMessages()
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "failed to get all leaving messages",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, messages)
	}
}

func SelectCommentMessages(c *gin.Context) {
	var m models.Message
	err := c.ShouldBind(&m)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}

	messages, err := service.SelectCommentMessages(m.UserName, m.Id)
	if err != nil {
		c.JSON(200, gin.H{
			"MSG":        "Select Comment Message Failed...",
			"error info": fmt.Sprintf("%v", err),
		})
	} else {
		c.JSON(200, messages)
	}
}
