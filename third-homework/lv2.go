package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileStr := "./plan.txt"
	file, err := os.Create(fileStr)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("close file error")
			return
		}
	}(file)
	if err != nil {
		panic("open file failed")
		return
	}
	str:="I’m not afraid of difficulties and insist on learning programming"
	_, err = file.Write([]byte(str))
	if err != nil {
		panic("write error")
		return
	}
	dst:=make([]byte,len(str))
	file1, err := os.Open(fileStr)
	if err != nil {
		return
	}
	_, err = file1.Read(dst)
	if err != nil {
		if err==io.EOF {
			fmt.Println("read all:")
		}else {
			panic("read error")
			return
		}
	}
	fmt.Println(string(dst))
}
