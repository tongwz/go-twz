package main

import (
	_ "unsafe"
)
import _ "linkname/pk"

func main() {
	printMySelf("hello world!")
}

//go:linkname printMySelf linkname/pk.myPrint
func printMySelf(a string)
