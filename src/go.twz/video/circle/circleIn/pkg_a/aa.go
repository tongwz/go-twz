package pkg_a

import _ "unsafe"

type A struct {
}

func (a *A) MyPrint() {
	printBB()
}

//go:noescape
//go:linkname printBB circleIn/pkg_b.PrintBB
func printBB()
