package main

import (
	"fmt"
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

	timeoutOccurred := make(chan bool, 1)
	go func() {
		timeoutOccurred <- WaitTimeout(&wg, time.Microsecond*5000)
	}()
	close(c)
	if <-timeoutOccurred {
		fmt.Println("timeout exit")
	} else {
		fmt.Println("Main: WaitGroup completed")
	}
	time.Sleep(time.Second * 6)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
	ch := make(chan bool)

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
