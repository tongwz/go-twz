package pkgB

import (
	"circle/pkgC"
	"fmt"
)

type B struct {
	C *pkgC.CS
}

func (b *B) MyPrint() {
	fmt.Println("我打印了 B")
}
