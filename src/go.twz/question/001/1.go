package main

import (
	"fmt"
)

func main() {
	// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case _, ok := <-ch1:
				if ok {
					fmt.Print(i)
					i++
					fmt.Print(i)
					i++
					ch2 <- true
				}
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case _, ok := <-ch2:
				if ok {
					if j > 'Z' {
						close(ch1)
						close(ch2)
						ch3 <- true
						return
					}
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					ch1 <- true
				}
			}
		}
	}()

	ch1 <- true

	<-ch3

}
