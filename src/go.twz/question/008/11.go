package main

import (
	"fmt"
)

type Peoplex interface {
	Show()
}

type Studentx struct{}

func (stu *Studentx) Show() {

}

func live() *Studentx {
	var stu *Studentx
	return stu
}

func main() {
	tmp := live()
	if tmp == nil {
		fmt.Println("AAAAAAA", tmp)
	} else {
		fmt.Println("BBBBBBB", tmp)
	}
}
