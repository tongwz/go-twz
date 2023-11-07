package main

import (
	"sync"
	"time"
)

// 总结：在程序
func main() {
	go LockFunc()
	time.Sleep(time.Second * 3)
}

func LockFunc() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Lock() // 确保即使发生 panic 也会执行 Unlock()

	// 在这里编写可能导致 panic 的代码
	panic("Something went wrong!")
}
