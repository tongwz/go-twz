package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000000; i++ {
		go addDoneFunc(&wg)
		go waitFunc(&wg)
	}

	wg.Wait()
}

func addDoneFunc(wg *sync.WaitGroup) {
	wg.Add(1)
	wg.Done()
}

func waitFunc(wg *sync.WaitGroup) {
	wg.Wait()
}
