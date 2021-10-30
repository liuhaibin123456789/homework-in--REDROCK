package main

import (
	"fmt"
)

func main() {
	channel3:=make(chan int)
	channel4:=make(chan bool)
	go func(channel chan int) {
		for i := 1; i < 101; i++ {
			<-channel
			if i%2==1 {
				fmt.Println(i)
			}
		}
	}(channel3)
	go func(channel chan int) {
		for i := 2; i < 101; i++ {
			channel<-i
			if i%2==0 {
				fmt.Println(i)
			}
			if i==100 {
				channel4<-true
			}
		}
	}(channel3)
	<-channel4
}
