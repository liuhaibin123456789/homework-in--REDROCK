package main

import (
	"fmt"
)
var(
	item int
)
func main() {
	//显示菜单页面
	fmt.Println("-----------1.登录-----------")
	fmt.Println("-----------2.注册-----------")
	fmt.Println("-----------3.退出-----------")
	fmt.Println("请选择标号：")
	_, err := fmt.Scanln(&item)
	if err != nil {
		panic(err)
	}
	switch item {
	case 1:
		err := login()
		if err != nil {
			panic(err)
		}
		//实现登录逻辑
	case 2:
		err := register()
		if err != nil {
			panic(err)
		}
		//实现注册逻辑
	case 3:
	default:
	}
}