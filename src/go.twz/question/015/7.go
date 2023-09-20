package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		fmt.Printf("我走到这里了2 %p \n", ch)
		ch <- 1
		fmt.Printf("结束1 \n")
	}()
	go func(ch chan int) {
		fmt.Printf("我走到这里了 %p \n", ch)
		time.Sleep(time.Second)
		<-ch
		fmt.Printf("结束2 \n")
	}(ch)

	var myMap map[string]string = nil
	fmt.Printf("打印没有make的map 值是：%s \n", myMap["基尼太美"])
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d \n", runtime.NumGoroutine())
	}
}
