package main

import "fmt"

type MyBook interface {
	FmtBook()
}

type Book struct {
	name string
}

func (b *Book) FmtBook() {
	fmt.Printf("我们在这里打印信息 \n")
	// fmt.Printf("我们在这里打印书名：%s \n", b.name)
}

func setOneBook() *Book {
	return nil
}

func FindFoo(a interface{}) {
	if a == nil {
		fmt.Printf("a == nil %#v ", a)
	} else {
		fmt.Printf("a != nil %#v ", a)
	}
}

func main() {
	var aBook MyBook
	aBook = setOneBook()
	if aBook == nil {
		fmt.Printf("我们的interface 返回值：%#v == nil \n", aBook)
	} else {
		fmt.Printf("我们的interface 返回值：%#v != nil \n", aBook)
		aBook.FmtBook()
	}

	var p *int = nil

	FindFoo(p)
}
