package main

import "fmt"

func main() {
	var slice1 []int
	slice2 := []int{}
	slice1 = append(slice1, 1)
	fmt.Printf("我们的slice1：%#v, %p, %d \n", slice1, slice1, len(slice1))
	fmt.Printf("我们的slice2：%#v, %p, %d \n", slice2, slice2, len(slice2))
}
