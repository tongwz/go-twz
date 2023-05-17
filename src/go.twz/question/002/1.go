package main

import "fmt"

func main() {
	var s = "²·`12345679ASDFGHJKLQWEasdfghjkl"
	if FindRepeat(s) {
		fmt.Printf("这个字符串不重复")
	} else {
		fmt.Printf("这个字符串重复的")
	}

}

func FindRepeat(s string) bool {
	if len(s) > 3000 {
		return false
	}

	for key, val := range s {
		if val > 256 {
			return false
		}
		for _, vv := range s[key+1:] {
			if vv == val {
				fmt.Println(string(val))
				return false
			}
		}

	}
	return true
}
