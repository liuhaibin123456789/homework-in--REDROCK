package main

import (
	"fmt"
)

const (
	rightAccount string="lhb123"
	rightPassword string="123456"
)
func check(inputAccount, inputPassword string) bool {
	var isEqual bool=false
	if inputAccount==rightAccount&&inputPassword==rightPassword {
		isEqual=true
	}
	return isEqual
}
func main() {
	var(
		inputAccount string
		inputPassword string
	)
	fmt.Println("请输入账号：")
	fmt.Scan(&inputAccount)
	fmt.Println("请输入密码：")
	fmt.Scan(&inputPassword)
	if check(inputAccount, inputPassword) {
		fmt.Println("登陆成功~")
	}else {
		fmt.Println("账号或密码错误！")
	}
}
