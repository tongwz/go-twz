package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "R10(FF)"
	// RLFLF

	nS := repeatStr(str)
	fmt.Printf("我们的str: %s", nS)

	x, y := xyFunc(nS)
	fmt.Printf("我们的值是 x：%d, y: %d", x, y)
}

func xyFunc(s string) (x, y int) {
	ss := []rune(s)
	f := "Y"
	for _, val := range ss {
		if val == 'R' {
			switch f {
			case "Y":
				f = "X"
			case "X":
				f = "-Y"
			case "-Y":
				f = "-X"
			case "-X":
				f = "Y"
			}
		} else if val == 'L' {
			switch f {
			case "Y":
				f = "-X"
			case "X":
				f = "Y"
			case "-Y":
				f = "X"
			case "-X":
				f = "-Y"
			}
		} else if val == 'F' {
			switch f {
			case "Y":
				y++
			case "X":
				x++
			case "-Y":
				y--
			case "-X":
				x--
			}
		} else if val == 'B' {
			switch f {
			case "Y":
				y--
			case "X":
				x--
			case "-Y":
				y++
			case "-X":
				x++
			}
		}
	}

	return x, y
}

func repeatStr(s string) (newStr string) {
	ss := []rune(s)
	l := len(ss)
	tmpNum := 0

	for i := 0; i < l; i++ {
		if isNumber(ss[i]) {
			tmpNumStr, tmpNumLen := repeatNum(ss[i:])
			fmt.Println(tmpNumStr, tmpNumLen)
			tmpNum, _ = strconv.Atoi(tmpNumStr)
			i += tmpNumLen - 1
			continue
		}
		if ss[i] == '(' {
			tmpSt, tmpStLen := repStr(ss[i+1:])
			newStr += strings.Repeat(tmpSt, tmpNum)
			i += tmpStLen + 1
		} else {
			newStr += string(ss[i])
		}

	}

	return newStr
}

func repeatNum(s []rune) (string, int) {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			return string(s[:i]), i
		}
	}
	return "", 0
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func repStr(s []rune) (string, int) {
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == ')' {
			return string(s[:i]), i
		}
	}
	return "", 0
}
