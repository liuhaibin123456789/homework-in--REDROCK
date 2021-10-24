package main

import (
	"fmt"
	"math/rand"
	"time"
)
func sort(slice []int) {
	for i:=0;i<len(slice)-1;i++ {
		for j:=0;j<len(slice)-1-i;j++ {
			if slice[j]<slice[j+1] {
				temp:=slice[j]
				slice[j]=slice[j+1]
				slice[j+1]=temp
			}
		}
	}
}
func main()  {
	numberSlice := make([]int, 10)
	for i := 0; i < 100; i++ {
		//设置随机种子
		rand.Seed(time.Now().UnixNano())
		//睡眠，防止程序过快，得到的随机数一致
		time.Sleep(time.Microsecond)
		//设置随机数
		randNumber := rand.Intn(100)
		numberSlice = append(numberSlice, randNumber)
	}
	sort(numberSlice)
	fmt.Println(numberSlice)
}
