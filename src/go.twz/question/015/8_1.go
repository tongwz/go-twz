package main

import "fmt"

func sendData(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func receiveData(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	receiveData(ch)

	// ch2 := make(<-chan int, 2)
	//
	// close(ch2)
}
