package main

import (
	"fmt"
)

var (
	x       int64
	channel1 chan int64
	channel2 chan bool
)

func add(channel chan int64) {
	for i := int64(0); i < 50000; i++ {
		<-channel
		x++
		channel1<-i
		if x == 100000 {
			channel2<-true
		}
	}
}
func main() {
	channel1=make(chan int64)
	channel2=make(chan bool)
	go add(channel1)
	go add(channel1)
	channel1<-0//启动因子，放在所有协程之后，若放在前面会导致主协程阻塞，从而通道死锁
	<-channel2
	fmt.Println(x)
}
