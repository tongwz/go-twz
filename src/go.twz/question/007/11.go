package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("Dropping mic")

	// 这里会让渡出当前goroutine的执行权，但是它并没有挂起当前的goroutine，因此她后续还是会自动恢复执行；
	// Yield execution to force executing other goroutines
	start := time.Now()
	runtime.Gosched()

	runtime.GC()
	cost := time.Since(start)

	fmt.Println("Done", cost)
}
