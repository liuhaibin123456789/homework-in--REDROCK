package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/models"
	"message-board/service"
)

func Login(c *gin.Context) {
	var u models.User
	err := c.ShouldBind(&u) //反射绑定json及form
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}

	res, err := service.LoginByPwd(u.Account, u.Password)
	if res {
		c.JSON(200, gin.H{
			"MSG": u.Account + "，welcome to message-board~",
		})
	} else {
		c.JSON(200, gin.H{
			"MSG":        "login failed, please check the account or password...",
			"error info": fmt.Sprintf("%v", err),
		})
	}
}

func Register(c *gin.Context) {
	var u models.User
	err := c.ShouldBind(&u) //反射绑定json
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	res, err := service.Register(u)
	if res {
		c.JSON(200, gin.H{
			"MSG": u.Account + ",register is OK in message-board~",
		})
	} else {
		c.JSON(200, gin.H{
			"error info": fmt.Sprintf("%v", err),
		})
	}
}

func UpdatePwd(c *gin.Context) {
	var u models.User
	err := c.ShouldBind(&u)
	if err != nil {
		fmt.Println("ShouldBind failed...", err)
		return
	}
	//获取旧密码
	oldPwd := c.PostForm("oldPwd")
	//核查账户及旧密码存在性
	if res1, err1 := service.CheckUser(u.Account, oldPwd); res1 && err1 == nil {
		err2 := service.UpdatePwd(u.Account, u.Password)
		if err2 == nil {
			c.JSON(200, gin.H{
				"MSG": "update succeed...",
			})
		} else {
			c.JSON(200, gin.H{
				"MSG":        "update failed but the user info is correct...",
				"error info": fmt.Sprintf("%v", err2),
			})
		}
	} else {
		c.JSON(200, gin.H{
			"MSG":        "the old user info is wrong...",
			"error info": fmt.Sprintf("%v", err1),
		})
	}
}
