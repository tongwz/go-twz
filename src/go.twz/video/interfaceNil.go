package main

import (
	"fmt"
	"unsafe"
)

type MyBook interface {
	FmtBook()
}

type Book struct {
	name string
}

func (b *Book) FmtBook() {
	fmt.Printf("我们在这里打印信息 \n")
}

func setOneBook() MyBook {
	return nil
}

type FindFoo interface {
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

	var p FindFoo
	var q *int = nil
	p = q

	length := unsafe.Sizeof(p)

	if p == nil {
		fmt.Printf("p == nil %#v  \n", p)
	} else {
		fmt.Printf("p != nil %#v  \n", p)
	}

	fmt.Printf("interface长度是：%d \n", length)

	var b MyBook
	length2 := unsafe.Sizeof(b)
	fmt.Printf("interface长度是：%d \n", length2)
	// b.FmtBook()

}

// go tool compile -S -p twzz -importcfg importcfg.txt interfaceNil.go
