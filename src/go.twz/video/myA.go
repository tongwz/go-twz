package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 2>&1
func main() {
	for {
		fmt.Println("标准输出~")
		io.WriteString(os.Stderr, "标准错误输出~ \n")
		time.Sleep(time.Second * 1)
	}
}
