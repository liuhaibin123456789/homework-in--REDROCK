package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan interface{})
	var group sync.WaitGroup
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func(x int) {
			defer group.Done()
			<-channel
			fmt.Println(x)
		}(i)
	}
	close(channel)
	group.Wait()
	fmt.Println("over!!!")

}
