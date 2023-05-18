package main

import "fmt"

func main() {
	str := "sdfsf我爱helloworld"
	reversString(&str)
	fmt.Printf("新字符串为：%s", str)
}

func reversString(s *string) bool {
	strS := []int32(*s)

	length := len(strS)

	if length > 5000 {
		return false
	}
	for key, _ := range strS {
		if key > length/2 {
			break
		}
		strS[key], strS[length-key-1] = strS[length-key-1], strS[key]
	}
	ss := string(strS)
	*s = ss
	return true
}
