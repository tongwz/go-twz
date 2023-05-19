package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "112中国34567qwe我爱你rtyY"
	s2 := "123我爱你中国Y4567qwerty1"

	if checkStr(s1, s2) {
		fmt.Println("我们的两个字符串重排之后可以相等")
	} else {
		fmt.Println("我们的两个字符串重排之后不相等")
	}

}

func checkStr(s1, s2 string) bool {
	ss1 := []rune(s1)
	ss2 := []rune(s2)

	if len(ss1) != len(ss2) {
		return false
	}

	if len(ss1) > 5000 || len(ss2) > 5000 {
		return false
	}

	for _, val := range ss1 {
		if strings.Count(s1, string(val)) != strings.Count(s2, string(val)) {
			return false
		}
	}

	return true
}
