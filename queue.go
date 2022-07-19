// @author cold bin
// @date 2022/7/19

package main

import (
	"errors"
	"fmt"
	"reflect"
)

type TSlice struct {
	I        []interface{}
	Remember reflect.Kind //记住类型：以第一次push的元素类型为准
	Front    int          //队头
	Rear     int          //队尾
}

func NewTSlice() TSlice {
	return TSlice{I: make([]interface{}, 0), Rear: -1, Front: -1, Remember: 0}
}

func (t *TSlice) CheckType(ele interface{}) bool {
	newVal := reflect.ValueOf(ele)

	if newVal.Type().Kind() != t.Remember {
		return false
	}
	return true
}

func (t *TSlice) isEmpty() bool {
	return t.Front == t.Rear
}

func (t *TSlice) push(ele interface{}) error {
	if t.isEmpty() && t.Remember == 0 {
		t.Remember = reflect.ValueOf(ele).Type().Kind()
		t.I = append(t.I, ele)
		return nil
	}
	if t.CheckType(ele) {
		t.Rear++
		t.I = append(t.I, ele)
		return nil
	}
	return errors.New("not the type,must be type of the first ele")
}

func (t *TSlice) pop() (interface{}, error) {
	if t.isEmpty() {
		return nil, errors.New("empty, no ele")
	}
	t.Front++
	ele := t.I[t.Front]
	return ele, nil
}

func (t *TSlice) peek() (interface{}, error) {
	if t.isEmpty() {
		return nil, errors.New("empty, no ele")
	}

	ele := t.I[t.Front+1]
	return ele, nil
}

func (t *TSlice) length() int {
	return t.Rear - t.Front + 1
}

func main() {
	tSlice := NewTSlice()

	err := tSlice.push("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tSlice.push("2")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tSlice.push("3")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tSlice.push("4")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tSlice.push("5")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tSlice.I)
	fmt.Println(tSlice.peek())
	fmt.Println(tSlice.pop())
	fmt.Println(tSlice.peek())

	fmt.Println(tSlice.length())
}
