package main


import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func register() error {
	var account,password string
	//获取注册账号密码
	fmt.Println("TIPS:账号20位，密码不超过六位")
	fmt.Println("请输入账号：")
	for true {
		_, err := fmt.Scanln(&account)
		if err != nil {
			return  err
		}
		if len(account)>20 {
			fmt.Println("账号字段过长，请重新输入：")
			continue
		}else {
			break
		}
	}
	fmt.Println("请输入密码：")
	for true {
		_, err1 := fmt.Scanln(&password)
		if err1!=nil {
			return err1
		}
		if len(password) > 6 {
			fmt.Println("密码输入错误，请重新输入：")
			continue
		}else {
			break
		}
	}
	var isRegister string
	fmt.Println("是否保存？ （y/n）")
	_, err := fmt.Scanln(&isRegister)
	if err != nil {
		panic(err)
	}
	if strings.EqualFold("y",isRegister) {
		//将账号密码保存至json格式化，保存至指定文件夹
		map1:=map[string]string{
			"账号":account,
			"密码":password,
		}
		//序列化
		jsonSlice, err := json.Marshal(map1)
		if err != nil {
			panic(err)
		}
		//保存至文件夹
		writeFile, err := os.OpenFile("d:/user.data",os.O_EXCL|os.O_WRONLY, 0666)
		if err != nil {
			panic("writeFile open error")
		}
		defer func(writeFile *os.File) {
			err := writeFile.Close()
			if err != nil {
				panic(err)
			}
		}(writeFile)
		_, err = writeFile.Write(jsonSlice)
		if err != nil{
			panic(err)
		}
		//将英文逗号作为分隔符
		_, err = writeFile.WriteString("\n")
		if err != nil {
			return err
		}
		fmt.Println("密码保存成功~")
	}else {
		fmt.Println("你的账号密码没有保存~")
	}
	return nil
}
