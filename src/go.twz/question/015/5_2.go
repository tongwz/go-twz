package main

import "sync"

func main() {
	var mu = struct {
		count int
		sync.Mutex
	}{}

	mu.Lock()

	mu.Lock()
}
