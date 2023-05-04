package main

import "fmt"

type MyBook interface {
	SetBook() error
}

type Book struct {
}

func (b *Book) SetBook() error {
	return nil
}

func setOneBook() MyBook {
	return nil
}

func main() {
	var aBook MyBook
	aBook = setOneBook()
	if aBook == nil {
		fmt.Printf("我们的interface 返回值：%#v == nil \n", aBook)
	} else {
		fmt.Printf("我们的interface 返回值：%#v != nil \n", aBook)
	}
}
