package main

import (
	"sync"
	"time"
)

// 计数器 负数情况
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Done()
	}()
	wg.Wait()
}
