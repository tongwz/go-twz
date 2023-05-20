package main

import "fmt"

func main() {
	str := "sdjfsld    sdf"

	newStr, _ := changeMyStr(str)
	fmt.Printf("获取的结果：%s", newStr)

}

func changeMyStr(s string) (string, bool) {
	ss := []rune(s)
	if len(ss) > 1000 {
		return "长度不符合要求", false
	}

	ss1 := []rune{}
	ss2 := []rune("%20")

	for _, val := range ss {
		if val != ' ' && (val > 'Z' || val < 'A') && (val > 'z' || val < 'a') {
			return string(val) + "字符不符合要求", false
		}
		if val != ' ' {
			ss1 = append(ss1, val)
		} else {
			ss1 = append(ss1, ss2...)
		}
	}
	return string(ss1), true
}
