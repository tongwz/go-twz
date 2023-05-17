package main

import (
	"fmt"
)

func main() {
	s := "··²123456asdfgASDGGH"
	if IsRepeat(s) {
		fmt.Println("不重复")
	} else {
		fmt.Println("重复")
	}
}

func IsRepeat(s string) bool {
	if len(s) > 3000 {
		return false
	}
	for key, val := range s {
		fmt.Println(val)
		if val > 256 {
			return false
		}
		for kk, vv := range s {
			if vv == val && key != kk {
				fmt.Println(string(vv))
				return false
			}
		}
	}
	return true
}
