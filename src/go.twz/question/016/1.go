package main

import (
	"container/list"
	"fmt"
)

func main() {
	var myList list.List

	myList.PushFront("1")
	myList.PushFront("2")
	myList.PushFront("3")

	myList.PushBack("4")

	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Printf("我的值是：%s \n", i.Value)
	}
}
