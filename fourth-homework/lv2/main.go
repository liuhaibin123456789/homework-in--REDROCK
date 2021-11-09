package main

import (
	"fmt"
	"time"
)
/*
     ● 在任意分钟的2秒时，输出"谁能比我卷！"
     ● 在任意分钟的3秒时，输出"早八算什么，早六才是吾辈应起之时"
     ● 自程序运行时起，每过3秒输出"芜湖！起飞！"
*/
func main() {
	go oneTask()
	go otherTask1()
	go otherTask2()
	time.Sleep(time.Minute)
}
//负责计时
func oneTask()  {
	if time.Now().Second() != 2 {
		time.Sleep(time.Second)
	}
	ticker := time.NewTicker(time.Second * 2)
	for true {
		select {
		case <-ticker.C:
			fmt.Println("谁能比我卷！")
		}
	}
}

func otherTask1() {
	if time.Now().Second() != 3 {
		time.Sleep(time.Second)
	}
	ticker := time.NewTicker(time.Second * 6)
	for true {
		select {
		case <-ticker.C:
			fmt.Println("早八算什么，早六才是吾辈应起之时")
		}
	}
}
func otherTask2() {
	ticker := time.NewTicker(time.Second * 3)
	for true {
		select {
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}
}