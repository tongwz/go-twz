package main

import "fmt"

func main() {
	var s string
	zhoujielun(s)
	var ss interface{}
	zhoujielun(ss)
}

type student struct {
	Name string
}

func zhoujielun(v interface{}) {
	switch msg := v.(type) {
	case *student:
		_ = msg.Name
		fmt.Println("是*student类型")
	case student:
		_ = msg.Name
		fmt.Println("是student类型")
	default:
		fmt.Printf("类型是：%T", msg)
	}
}
