package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) IncrementAge() {
	p.Age++
}

func main() {
	p := Person{Name: "Alice"}
	p.IncrementAge()
	fmt.Println(p.Age) // Output: 1
}
