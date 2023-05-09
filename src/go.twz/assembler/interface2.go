package assembler

import "fmt"

type Inter2 interface {
	My()
}

type InterS2 struct {
}

func (i *InterS2) My() {
	fmt.Println("test")
}

func p3e(i *InterS2) Inter2 {
	return i
}
