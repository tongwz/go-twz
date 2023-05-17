package main

import _ "unsafe"

//go:linkname printnl runtime.printnl
func printnl()

//go:linkname printstring runtime.printstring
func printstring(s string)

func printlnn(s string) {
	println(s)
}

var helloworld = "你好，世界"

func main()
