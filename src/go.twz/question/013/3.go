package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		go SetFunc()
	}
	time.Sleep(time.Second * 3)
	numGoroutines := runtime.NumGoroutine()
	fmt.Println("当前程序中有", numGoroutines, "个活跃的 goroutine")

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v iB \n", m.Alloc)

	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v iB \n", m.Alloc)

	time.Sleep(time.Second * 3)
}

func SetFunc() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			// fmt.Println(num)
		}(i, c)
	}
	numGoroutines := runtime.NumGoroutine()
	fmt.Println("当前程序中有", numGoroutines, "个活跃的 goroutine")

	if WaitTimeout(&wg, time.Second*2) {
		close(c)
		// fmt.Println("timeout exit")
	}
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	// ch := make(chan bool)
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
