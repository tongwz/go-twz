package main

import _ "unsafe"
import _ "linkname/pk"

func main() {
	myPrint("tongWz")
}

//go:linkname myPrint linkname/pk.printMyInfo
func myPrint(a string)
