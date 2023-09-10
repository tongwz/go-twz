package main

import (
	"fmt"
	"sync"
	"time"
)

type MyPoolStr struct {
	Name string
}

var myPoolStr = sync.Pool{
	New: func() interface{} {
		return new(MyPoolStr)
	}}

func main() {

	processRequest3("基")  // 1KiB
	processRequest3("尼")  // 1KiB
	processRequest3("太美") // 1KiB

	// time.Sleep(time.Millisecond)
	bb := myPoolStr.Get().(*MyPoolStr)

	fmt.Printf("b的指针是2：%p , name = %s \n", bb, bb.Name)

	bb = myPoolStr.Get().(*MyPoolStr)

	fmt.Printf("b的指针是2：%p , name = %s \n", bb, bb.Name)

}
func processRequest3(name string) {
	b := myPoolStr.Get().(*MyPoolStr)
	b.Name = name
	fmt.Printf("b的指针是：%p, name= %s \n", b, b.Name)
	myPoolStr.Put(b)
	time.Sleep(1 * time.Millisecond)
}
