package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		timer := time.NewTicker(time.Second)

		for {
			select {
			case <-timer.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Printf("我们捕获了异常： %s \n", err)
						}
					}()
					proc()
				}()
			}
		}

	}()

	select {}
}

func proc() {

	panic("ok")
}
