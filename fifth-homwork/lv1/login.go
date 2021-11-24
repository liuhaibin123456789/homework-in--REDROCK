package main

//使用cookie实现登录状态，写个中间件控制访问资源的权限
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//loginMiddleware 实现对登录的账号及密码的检查
func loginMiddleware(c *gin.Context)  {
	//获取cookie
	cookie, err := c.Cookie("gin_cookie")
	//中间件逻辑
	defer func(c *gin.Context) {
		if err!=nil {
			//获取失败
			c.String(http.StatusForbidden,"not find cookie...")
		}else {
			//获取成功
			c.Set("cookie",cookie)
			c.Next()
		}
	}(c)
}
func main() {
	//创建路由
	router := gin.Default()
	//web服务
	router.POST("/login",func(c *gin.Context) {
		//获取登录参数
		user := c.PostForm("user")
		pwd := c.PostForm("pwd")
		if user=="lhb"&&pwd=="123" {
			//校验正确，保存cookie
			c.SetCookie("gin_cookie",user,3306,"/","",false,false)
			c.JSON(200,gin.H{
				user:"欢迎登陆！",
			})
		}else {
			c.String(http.StatusForbidden,"账号或密码错误")
		}
	})
	router.GET("/webServer1",loginMiddleware,func(c *gin.Context) {
		get, exists := c.Get("cookie")
		if !exists {
			c.String(403,"not exist...")
			return
		}else {
			c.JSON(200,gin.H{
				get.(string):" 欢迎回来!",
			})
		}
	})
	//监听端口
	router.Run(":8081")
}