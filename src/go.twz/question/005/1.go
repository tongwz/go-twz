package main

import (
	"fmt"
)

func main() {
	s1 := "abcdefg LKLJ sdf "

	fmt.Println(' ')

	s2, isReal := changeStr(s1)
	if !isReal {
		fmt.Printf(s2)
	} else {
		fmt.Println(s2)
	}
}

func changeStr(s string) (string, bool) {
	ss1 := []rune(s)
	if len(ss1) > 1000 {
		return "长度不符合要求", false
	}
	ss2 := []rune{}
	sliceT := []rune("%20")

	for _, val := range ss1 {
		if val != ' ' && (val > 'Z' || val < 'A') && (val > 'z' || val < 'a') {
			return string(val) + "不符合要求", false
		}

		if val == ' ' {
			ss2 = append(ss2, sliceT...)
		} else {
			ss2 = append(ss2, val)
		}
	}
	return string(ss2), true
}
