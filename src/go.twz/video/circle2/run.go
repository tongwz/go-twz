package main

import (
	"circle2/pkgAA"
	"time"
)

func main() {
	a := new(pkgAA.MyA)

	a.MyPrint()

	time.Sleep(time.Second)
}
