package main

import (
	"fmt"
	"time"
)

func main() {
	var myC = make(chan int)
	var myC2 = make(chan int, 1)

	myC <- 1
	myC2 <- 2

	// channel 长度 为0 如果我们不进行预先有协程进行读取，那就直接阻塞当前 线程
	go myRun(myC)
	go myRun(myC2)
	time.Sleep(time.Second * 2)
}

func myRun(ch chan int) {
	myVal := <-ch
	fmt.Println("获取到值", myVal)
}
