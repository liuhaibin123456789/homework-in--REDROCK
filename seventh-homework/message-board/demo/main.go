package main

import (
	"fmt"
	"message-board/cmd"
	"message-board/tool"
)

//项目入口
func main() {
	//初始化数据库连接
	err := tool.InitSQL()
	if err != nil {
		fmt.Println(err)
		return
	}
	//启动路由
	cmd.AccessUrl()
}
