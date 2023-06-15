package main

import (
	"fmt"
	"sync"
)

type query1 func(string) string

func exec1(name string, vs ...query1) string {
	var wg sync.WaitGroup
	ch := make(chan string)
	for _, v := range vs {
		wg.Add(1)
		go func(v query1) {
			defer wg.Done()
			ch <- v(name)
		}(v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return <-ch
}

func main() {
	ret := exec1("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
