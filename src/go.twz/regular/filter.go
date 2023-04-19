package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "callFile/4y7il5/7ACF6B9576834E168BEC68A313CFB5DC/2022/1110/15616130872_20221110092707_180000.mp3"

	re := strings.Split(str, "/")
	res := re[len(re)-1]

	fmt.Printf("值是：%#v \n", res)
}
