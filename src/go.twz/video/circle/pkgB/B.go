package pkgB

import (
	"circle/pkgA"
	"fmt"
)

type B struct {
	C *pkgA.CS
}

func (b *B) MyPrint() {
	fmt.Println("我打印了 B")
}
