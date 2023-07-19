package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = iota
	age  = 100
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
