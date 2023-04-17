package main

import (
	"fmt"
	"regexp"
)

func main() {
	regularRule := "^1[1-9]{1}\\d{9}$"

	reg := regexp.MustCompile(regularRule)
	phone := "13355556666"
	if reg.MatchString(phone) {
		fmt.Println("是手机号")
	} else {
		fmt.Println("不是手机号")
	}
}
