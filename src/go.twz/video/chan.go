package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	run(ch)
	time.Sleep(time.Second * 1)
}

func run(ch chan int) {
	myVal := <-ch
	fmt.Println("获取到值", myVal)
}
