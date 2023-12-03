package main

import (
	"fmt"
	"learn01/pkg"
)

// import pkg "int.go"

func main() {
	fmt.Printf("我的名字是：%s", "吴彦祖")
	println(pkg.Id)
	println(pkg.Name)
	pkg.NameData[0] = '?'
	println(pkg.Name)
	println(pkg.Name2)
}
