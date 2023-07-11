package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		select {
		case re := <-ch:
			fmt.Printf("获取的值是：%v \n", re)
		}
	}()

	close(ch)
	close(ch)
	time.Sleep(time.Second)
}
