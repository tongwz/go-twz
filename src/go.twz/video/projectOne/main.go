package main

import "fmt"

func main() {
	fmt.Printf("我字对应的 unicode是：%d \n", '我')
	fmt.Printf("我字对应的 unicode是：%b \n", '我')

	str := "我"
	slice := []byte(str)
	fmt.Printf("我字对应的 utf8是：%b \n", slice)
}
