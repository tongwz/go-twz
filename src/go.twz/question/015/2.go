package main

import (
	"fmt"
	"sync"
	"time"
)

var mu2 sync.RWMutex
var count int

func main() {
	go A2()
	time.Sleep(2 * time.Second)
	mu2.Lock()
	defer mu2.Unlock()
	count++
	fmt.Println(count)
}
func A2() {
	mu2.RLock()
	defer mu2.RUnlock()
	B2()
}
func B2() {
	time.Sleep(5 * time.Second)
	C2()
}
func C2() {
	mu2.RLock()
	defer mu2.RUnlock()
}
