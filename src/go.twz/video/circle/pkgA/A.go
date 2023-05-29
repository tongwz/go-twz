package pkgA

import "circle/pkgB"

// type MyInterface interface {
// 	MyPrint()
// }

type A struct {
	C *CS
	B *pkgB.B
}

func (a *A) MyPrint() {
	a.B = new(pkgB.B)
	a.B.MyPrint()
}
