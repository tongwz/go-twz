package main

import (
	"circle/pkgA"
	"circle/pkgB"
	"time"
)

func main() {
	a := new(pkgA.A)

	a.B = new(pkgB.B)

	a.MyPrint()

	time.Sleep(time.Second * 2)
}
