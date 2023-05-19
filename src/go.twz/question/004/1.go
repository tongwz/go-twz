package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "1uqwerty我ipp1234466789"
	s2 := "qwerty我ipp1234466u7891"
	if IsSame(s1, s2) {
		fmt.Println("字符串可以排序后可以一致")
	} else {
		fmt.Println("字符串可以排序后不可以一致")
	}

	ss1, _ := SortStr([]rune(s1))
	fmt.Printf("我们获取到的排序后的值是 ： %s \n", string(ss1))

	ss2, _ := SortStr([]rune(s2))
	fmt.Printf("我们获取到的排序后的值是 ： %s \n", string(ss2))
	if string(ss1) == string(ss2) {
		fmt.Println("字符串可以排序后可以一致")
	} else {
		fmt.Println("字符串可以排序后不可以一致")
	}

}

func IsSame(s1, s2 string) bool {
	ss1 := []rune(s1)
	ss2 := []rune(s2)
	l1 := len(ss1)
	l2 := len(ss2)
	if l1 != l2 || l1 > 5000 || l2 > 5000 {
		return false
	}

	for _, v := range ss1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}

// 147389

func SortStr(ss []rune) ([]rune, bool) {
	l := len(ss)
	isChange := false
	for key, _ := range ss {
		if l-1 > key && ss[key] > ss[key+1] {
			ss[key], ss[key+1] = ss[key+1], ss[key]
			isChange = true
		}
	}
	if !isChange {
		return ss, false
	} else {
		return SortStr(ss)
	}
}
