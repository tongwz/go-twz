package main

import (
	"fmt"
)

func main() {
	fmt.Printf("打印类型：%v, %T \n", [...]string{"1"}, [...]string{"1"})
	fmt.Printf("打印类型：%v, %T \n", []string{"1"}, []string{"1"})
	fmt.Println([...]string{"1"} == [...]string{"1"})
	// fmt.Println([]string{"1"} == []string{"1"})
}
