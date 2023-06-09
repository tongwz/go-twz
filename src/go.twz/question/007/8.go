package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}

	go func() {
		// for a := range abc {
		// 	fmt.Println("a: ", a)
		// }

		for {
			select {
			case a, ok := <-abc:
				if ok {
					fmt.Println("A: ", a)
				} else {
					fmt.Println("AA: ", a)
				}
			}
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Microsecond * 1)
}
