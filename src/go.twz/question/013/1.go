package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}
	numGoroutines := runtime.NumGoroutine()
	fmt.Println("当前程序中有", numGoroutines, "个活跃的 goroutine")

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second)
	numGoroutines = runtime.NumGoroutine()
	fmt.Println("当前程序中有", numGoroutines, "个活跃的 goroutine")
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool, 1)

	timer := time.NewTimer(timeout)
	go func() {
		select {
		case <-timer.C:
			ch <- true
		}
	}()

	go func() {
		wg.Wait()
		ch <- false
	}()

	return <-ch
}
