package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer_call()
	defer_call()
}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
	fmt.Println(21)
}

// 而函数的退出顺序是后进先出的。所以，最后一个 defer 函数会最先执行，而第一个 defer 函数会最后执行。这种执行方式可以确保资源的正确释放，例如文件句柄、数据库连接等资源的关闭。
