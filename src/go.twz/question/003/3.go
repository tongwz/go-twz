package main

import "fmt"

func main() {
	str := "我爱你中国1234"
	reverseStr(&str)
	fmt.Printf("我们打印的结果是：%s \n", str)
}

func reverseStr(s *string) bool {
	strS := []rune(*s)

	length := len(strS)

	if length > 5000 {
		return false
	}

	for key, _ := range strS {
		if key >= length/2 {
			break
		}
		strS[key], strS[length-key-1] = strS[length-key-1], strS[key]
	}
	*s = string(strS)

	return true
}
