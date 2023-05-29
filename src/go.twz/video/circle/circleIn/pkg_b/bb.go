package pkg_b

import (
	"circleIn/pkg_a"
	"fmt"
)

type B struct {
	C *pkg_a.CS
}

func (b *B) MyPrint() {
	fmt.Println("我打印了 B")
}

func PrintBB() {
	fmt.Println("我打印了 PrintBB")
}
