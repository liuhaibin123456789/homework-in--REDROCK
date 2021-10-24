package main

import "fmt"

func sort(slice []int32) {
	for i:=0;i<len(slice)-1;i++ {
		for j:=0;j<len(slice)-1-i;j++ {
			if slice[j]>slice[j+1] {
				temp:=slice[j]
				slice[j]=slice[j+1]
				slice[j+1]=temp
			}
		}
	}
}
func main()  {
	numberSlice:=make([]int32,10)
	fmt.Println("enter elements of the array:")
	for i := 0; i < len(numberSlice); i++ {
		fmt.Scan(&numberSlice[i])
	}
	sort(numberSlice)
	fmt.Println(numberSlice)
}
