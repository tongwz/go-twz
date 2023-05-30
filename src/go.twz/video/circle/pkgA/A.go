package pkgA

import (
	"circle/pkgB"
	"circle/pkgC"
)

type MyInterface interface {
	MyPrint()
}

type A struct {
	C *pkgC.CS
	B *pkgB.B
}

func (a *A) MyPrint() {
	a.B.MyPrint()
}
