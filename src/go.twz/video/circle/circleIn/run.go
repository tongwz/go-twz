package main

import "circleIn/pkg_a"
import _ "circleIn/pkg_b"

func main() {
	a := new(pkg_a.A)

	a.MyPrint()
}
