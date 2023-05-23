package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "R10(RF)F"
	newStr := splitString(str)
	fmt.Println(newStr)
	x, y := operation(newStr)
	fmt.Println(x, y)
}

/**
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。 问最后机器人的坐标是多少？
*/

func splitString(str string) string {
	slice := []rune(str)
	l := len(slice)
	newStr := ""

	// 遇到的数字
	tmpNum := 0
	// 括号内容
	// tmpComm := ""
	for i := 0; i < l; i++ {
		if isNum(slice[i]) {
			numStr, numStrlen := RepeatNum(slice[i:])
			fmt.Println(slice[i:], numStr, numStrlen)
			// 将unicode转化成数字
			tmpNum, _ = strconv.Atoi(numStr)
			i += numStrlen - 1
			continue
		}
		if slice[i] == '(' {
			tmpComm, lenTmpComm := RepeatStr(slice[i+1:])
			i += lenTmpComm + 1

			newStr += strings.Repeat(tmpComm, tmpNum)
		} else {
			newStr += string(slice[i])
		}
	}
	return newStr

}

// 判断是否是数字
func isNum(ru rune) bool {
	return '0' <= ru && ru <= '9'
}

func RepeatNum(st []rune) (string, int) {
	l := len(st)
	for i := 0; i < l; i++ {
		if st[i] == '(' {
			return string(st[:i]), i
		}
	}
	return "", 0
}

// RepeatStr 返回 重复的str 和 str长度
func RepeatStr(st []rune) (string, int) {
	l := len(st)
	for i := 0; i < l; i++ {
		if st[i] == ')' {
			return string(st[:i]), i + 1
		}
	}
	return "", 0
}

func operation(str string) (x, y int) {
	var dir = "+Y"
	sl := []rune(str)
	for _, val := range sl {
		if val == 'L' {
			switch dir {
			case "+Y":
				dir = "-X"
			case "-Y":
				dir = "+X"
			case "+X":
				dir = "+Y"
			case "-X":
				dir = "-Y"
			}
		} else if val == 'R' {
			switch dir {
			case "+Y":
				dir = "+X"
			case "-Y":
				dir = "-X"
			case "+X":
				dir = "-Y"
			case "-X":
				dir = "+Y"
			}
		} else if val == 'F' {
			switch dir {
			case "+Y":
				y++
			case "-Y":
				y--
			case "+X":
				x++
			case "-X":
				x--
			}
		} else if val == 'B' {
			switch dir {
			case "+Y":
				y--
			case "-Y":
				y++
			case "+X":
				x--
			case "-X":
				x++
			}
		}
	}
	return x, y
}
