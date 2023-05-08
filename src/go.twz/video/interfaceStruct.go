package main

import "fmt"

func main() {
	var abc interface{} = 1
	fmt.Println(abc)

	myBook := &OneBook{
		name:  "Go语言程序设计",
		price: 100.1,
	}

	var book GoBook = myBook

	fmt.Println(book)
}

type GoBook interface {
	Name()
	Price()
}

type OneBook struct {
	name  string
	price float64
}

func (m *OneBook) Name() {
	fmt.Printf("我们的书名是 ： %s \n", m.name)
}

func (m *OneBook) Price() {
	fmt.Printf("我们的书价格是 ： %f \n", m.price)
}
