package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)
/*
  每一功能模块都跑出panic，这样有问题时可以迅速找到问题所在
*/
func login() error {
	var account,password string
	//输入账号密码
	fmt.Println("账号:")
	_, err := fmt.Scanln(&account)
	if err != nil {
		panic(err)
	}
	fmt.Println("密码:")
	_, err = fmt.Scanln(&password)
	if err != nil {
		panic(err)
	}
	//读取正确账号密码
	str:="d:/user.data"
	file, err := os.OpenFile(str,os.O_RDONLY,0666)
	bufioFile:=bufio.NewReader(file)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	var b string
	b, err = bufioFile.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//反序列化
	var map1 map[string]string
	//map1=make(map[string]string)//自动创建
	err = json.Unmarshal([]byte(b), &map1)
	if err != nil {
		panic(err)
	}
	if map1["账号"] == account && map1["密码"] == password {
		fmt.Println("登陆成功~")
	}else {
		fmt.Println("登陆失败！")
	}
	return nil
}