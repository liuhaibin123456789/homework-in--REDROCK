package main
/*
旭宝是一个健忘的人，他想先注册一个账号，但是后面他忘记了密码，请使用密保这种形式将他的密码找回来并且进行更改。

● 登录+注册
● 可以通过密保（比如：你所就读的是哪一所高中，父亲叫什么名字什么的都行）来找回账号或者更改密码
*/

//定义全局数据库变量
import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

//定义全局数据库变量
var dba *sql.DB

//定义返回信息结构体，使反馈能更友好
type resultInfo struct {
	val string
	isSucceeded bool
	err error
}

//initSQL 初始化数据库，并校验是否可以连接数据库
func initSQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mydb5?charset=utf8"
	db, err := sql.Open("mysql", dsn)//校验连接语法是否正确
	if err != nil {
		return err
	}
	//defer db.Close()//延迟释放资源
	err = db.Ping()//检查是否能连接成功
	if err != nil {
		return err
	}
	dba=db
	return nil
}
//addUser 添加账户信息 question表示密保问题，answer表示密保问题答案，皆为可选项
func addUser(account, password, question,answer string) (err error) {
	sqlInsertStr:="insert into userInfo (`account`,`password`,`question`,`answer`)values (?,?,?,?)"
	_, err = dba.Exec(sqlInsertStr, account, password, question,answer)
	if err != nil {
		return err
	}
	return nil
}

//updateUser 用于修改用户密码 answer为外界校验
func updateUser(account,password string) resultInfo {
	//修改操作
	sqlUpdateStr:="update userInfo set password=? where account=?"
	_, err := dba.Exec(sqlUpdateStr, password, account)
	if err != nil {
		return resultInfo{
			err: err,
			isSucceeded: false,
			val: "password or account is false?",
		}
	}
	return resultInfo{
		isSucceeded: true,
		val: "password is updated successfully",
		err: nil,
	}
}

func isExist(account,password string) (res bool, err error) {
	sqlSelect := "select password from userInfo where account=?"
	var pwd string
	err = dba.QueryRow(sqlSelect, account).Scan(&pwd)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if pwd == password {
		return true,nil
	}
	return false,nil
}

//register web服务-register
func register(c *gin.Context)  {

	account := c.Query("account")
	password := c.Query("password")
	question := c.DefaultQuery("question","")
	answer := c.DefaultQuery("answer","")

	err := addUser(account, password, question, answer)
	if err != nil {
		c.JSON(403,gin.H{
			"MSG":"add user failed",
		})
		return
	}else {
		c.JSON(200,gin.H{
			"MSG":"register OK",
		})
	}
}

//login web服务-login
func login(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently,"https://www.bilibili.com/")
}

//modify 更改账户及密码
func modify(c *gin.Context) {
	//	/modify/:account/:password/:answer
	account := c.Param("account")
	password := c.Param("password")
	answer := c.Param("answer")
	isExist, err := isExist(account,password)
	if !isExist&&err == nil {
		//校验答案
		sqlQueueStr := "select answer from userInfo where account=?"
		var a string
		err := dba.QueryRow(sqlQueueStr, account).Scan(&a)
		if err != nil {
			return
		}
		if a==answer {
			//更改密码
			user := updateUser(account, password)
			c.JSON(200,gin.H{
				"MSG":user.val,
			})
		}else {
			c.JSON(200,gin.H{
				"MSG":"failed update user...the answer is wrong....",
			})
		}

	}else if !isExist&&err!=nil{
		c.JSON(200,gin.H{
			"MSG":"the account is not existent",
		})
	}else {
		c.JSON(200,gin.H{
			"MSG":"the password is existent",
		})
	}

}
//CheckMiddleware 中间件核查账户密码
func CheckMiddleware(c *gin.Context) {
	//初始化数据
	account := c.Query("account")
	password := c.Query("password")

	defer func(c *gin.Context) {
		res,_:=isExist(account,password)
		if res{
			c.Next()
		}else {
			c.Abort()
			c.JSON(200,gin.H{
				"MSG":"account or password is failed",
			})
		}
	}(c)
}

func main()  {
	err := initSQL()
	if err != nil {
		log.Println(err)
		return
	}
	//路由
	router := gin.Default()
	//注册服务
	router.POST("/register",register)
	//登录
	router.GET("/login",CheckMiddleware,login)
	//更改密码
	router.PUT("/modify/:account/:password/:answer",modify)
	//注册服务
	router.Run(":8081")
}
