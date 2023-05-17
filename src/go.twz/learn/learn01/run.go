package main

import "learn01/pkg"

// import pkg "int.go"

func main() {
	println(pkg.Id)
	println(pkg.Name)
	pkg.NameData[0] = '?'
	println(pkg.Name)
	println(pkg.Name2)
}
