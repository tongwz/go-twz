package pkgCC

import "circle2/pkgAA"

type MyC struct {
	a *pkgAA.MyA
}

func (m *MyC) MyPrint() {
	m.a.MyPrint()
}
