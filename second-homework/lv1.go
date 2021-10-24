package main

import "fmt"

/*
实现字符串倒序
*/
func main() {
	var old string
	fmt.Println("enter a string:")
	_, err := fmt.Scanf("%s",&old)
	if err != nil {
		panic("input error!")
	}
	//反转字符串
	oldBytes := []byte(old)
	for i := 0; i < len(oldBytes)/2; i++ {
		oldBytes[i],oldBytes[len(oldBytes)-1-i]=oldBytes[len(oldBytes)-1-i],oldBytes[i]
	}
	fmt.Println(string(oldBytes))
}
