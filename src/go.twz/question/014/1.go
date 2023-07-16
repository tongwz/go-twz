package main

import (
	"fmt"
)

type myType int

type mySt struct {
}

func main() {
	var x *string = nil
	var y []int = nil
	var z map[int]int = nil
	var a chan int = nil
	var b *myType = nil
	var c *mySt = nil
	if x == nil {
		yy := "default"
		x = &yy
	}
	fmt.Println(*x, y, z, a, b, c)
}
