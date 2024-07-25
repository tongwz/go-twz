package main

import (
	"fmt"
	"time"
)

// 这个应该是考察channel并发读取的问题 - 会有3个worker并发运行 但是 之后剩下2个 再被前两个消费完的 - 关闭了jobs这个channel 也会将剩余的值给读取完
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d \n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("work:%d end job:%d \n", id, j)
		results <- j * 2
	}
}
