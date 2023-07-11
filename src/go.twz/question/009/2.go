package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	wait := sync.WaitGroup{}
	wait.Add(5)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100)
		}
	}()

	go func() {
		for {
			select {
			case a, ok := <-ch:
				if ok {
					fmt.Printf("打印随机数： %d \n", a)
					wait.Done()
				}
			}
		}
	}()

	wait.Wait()
}
