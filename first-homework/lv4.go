package main

import (
	"fmt"
	"strings"
)

/*
你要发金币了，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后发出去了多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)
func main() {
	var sumCoin = 0
	for i := 0; i < len(users);i++ {
		distribution[users[i]]=0
	}
	for i := 0; i < len(users); i++ {
		dispatchCoin(users[i])
	}
	for key,value:=range distribution {
		fmt.Println(key,"得到了",value,"枚金币")
		sumCoin+=value
	}
	 fmt.Println("共有：",sumCoin,"枚金币")
}

func dispatchCoin(name string) {
	var coin = 0
	newName:=strings.ToLower(name)
	if strings.Contains(newName, "e") {
		coin+=1*strings.Count(newName,"e")
	}
	if strings.Contains(newName, "i"){
		coin+=2*strings.Count(newName,"i")
	}
	if strings.Contains(newName, "o"){
		coin+=3*strings.Count(newName,"o")
	}
	if strings.Contains(newName, "u") {
		coin+=4*strings.Count(newName,"u")
	}
	distribution[name]=coin
	coin=0
}