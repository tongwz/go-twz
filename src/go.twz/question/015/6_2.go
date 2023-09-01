package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var pool2 = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	}}

func main() {

	for i := 0; i < 15; i++ {
		go func() {
			for {
				processRequest2(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest2(size int) {
	b := pool2.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	fmt.Printf("我们的bytes.Buffer 设定长度是：%d \n", b.Cap())
	pool2.Put(b)
	time.Sleep(1 * time.Millisecond)
}
