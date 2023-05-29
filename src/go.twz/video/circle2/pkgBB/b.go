package pkgBB

import "circle2/pkgCC"

type MyB struct {
	c *pkgCC.MyC
}

func (m *MyB) MyPrint() {
	m.c = new(pkgCC.MyC)
	m.c.MyPrint()
}
