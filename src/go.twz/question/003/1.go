package main

import "fmt"

func main() {
	str := "122344556767dfgdhfghjghkjasfs"
	str2 := "我到此一游哈哈1"

	myReverseStr(&str)

	fmt.Printf("翻转后的值是str： %s \n", str)
	myReverseStr(&str2)
	fmt.Printf("翻转后的值是str2： %s \n", str2)
}

func myReverseStr(s *string) bool {
	strSli := []int32(*s)

	if len(strSli) > 5000 {
		return false
	}

	// 10
	length := len(strSli)

	for key, _ := range strSli {
		if key > (length / 2) {
			break
		}
		strSli[key], strSli[length-key-1] = strSli[length-key-1], strSli[key]
	}

	ss := string(strSli)
	*s = ss
	return true
}
