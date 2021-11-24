package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)
var filePath = "d:/users.data"

//saveUserInfo 保存信息至本地服务器文件
func saveUserInfo(account,password string) (err error) {
	//打开文件
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("can not open file....\nerr:",err)
		return err
	}
	//延迟关闭资源
	defer file.Close()
	//序列化数据
	userInfo:=make(map[string]string,1)
	userInfo[account]=password
	userBytes, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println(err)
		panic("userInfo marshal failed...")
		return err
	}
	//写入信息
	writer := bufio.NewWriter(file)
	_, err1 := writer.Write(userBytes)
	if err1 == nil {
		fmt.Println("账号密码写入成功！")
		fmt.Println(string(userBytes))
	}else {
		return err1
	}
	writer.Flush()
	//写入分隔符
	_, err = writer.WriteString("*")
	if err != nil {
		fmt.Println("':' write failed... ",err1)
		return err
	}
	writer.Flush()
	return nil
}
//getUserInfo 从本地服务器文件里拿出指定账户
func getUserInfo(account,password string) (isExist bool,err error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0)
	if err != nil {
		panic(err)
		return false, err
	}
	defer file.Close()
	//读取信息
	reader := bufio.NewReader(file)
	for true {
		//读取userInfo
		readString, err2 := reader.ReadString('*')
		//切除后缀
		readString = strings.TrimSuffix(readString, "*")
		//反序列化
		userInfo:=make(map[string]interface{},1)
		err = json.Unmarshal([]byte(readString), &userInfo)
		if err != nil {
			return false,errors.New("unmarshal failed")
		}
		//检验账户
		if password==userInfo[account] {
			return true,nil
		}
		//等待最后一行数据处理完毕,再退出
		if err2 != nil {
			if err2==io.EOF {
				break
			}else {
				return false,err2
			}
		}
	}
	return false,errors.New("not find the account")
}
func main() {
	//创建路由
	router := gin.Default()
	router.GET("/register", func(c *gin.Context) {
		//获取url账号密码参数
		account := c.Query("account")
		password := c.Query("password")
		//保存注册信息
		err := saveUserInfo(account, password)
		if err != nil {
			c.JSON(500,gin.H{
				"webSever":"ERROR!",
			})
		}else {
			c.SetCookie("cookie1",account,3600,"/","",false,false)
			c.JSON(200,gin.H{
				account: "register OK!",
			})
		}
	})
	router.GET("/login", func(c *gin.Context) {
		//获取当前登录用户信息
		account := c.Query("account")
		password := c.Query("password")
		//遍历本地服务器已有帐户
		isExist, err := getUserInfo(account,password)
		if err == nil&&isExist {
			//存在用户
			c.JSON(200,gin.H{
				account: "欢迎回来！",
			})
		}else {
			//不存在用户，再使用cookie默认登录
			cookie, err := c.Cookie("cookie1")
			if err != nil {
				c.String(403,"not find cookie...")
			}else {
				c.String(200,"%s, welcome back...",cookie)
			}
		}
	})
	//监听端口
	router.Run(":8082")
}