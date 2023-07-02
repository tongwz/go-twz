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

func live() Peoplex {
	var stu *Studentx
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
