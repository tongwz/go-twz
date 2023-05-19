package main

import "fmt"

func main() {
	s1 := "12中1国345啦啦啦67qwe我爱你rtyY"
	s2 := "123我爱你啦啦中国Y4567qw啦erty1"

	ss1, _ := sortStr([]rune(s1))
	ss2, _ := sortStr([]rune(s2))

	if string(ss1) == string(ss2) {
		fmt.Printf("我们两个字符串相等, %s", string(ss1))
	} else {
		fmt.Printf("我们两个字符串不相等, %s, %s", string(ss1), string(ss2))
	}
}

func sortStr(s []rune) ([]rune, bool) {
	l := len(s)
	isEnd := true
	for key, _ := range s {
		if l-1 > key && s[key] > s[key+1] {
			s[key], s[key+1] = s[key+1], s[key]
			isEnd = false
		}
	}
	if isEnd {
		return s, isEnd
	} else {
		return sortStr(s)
	}
}
