package pkgAA

import "circle2/pkgBB"

type MyA struct {
	b *pkgBB.MyB
}

func (m *MyA) MyPrint() {
	m.b = new(pkgBB.MyB)
	m.b.MyPrint()
}
