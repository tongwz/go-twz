package main

import (
	"fmt"
)

type StudentX struct {
	Age int
}

func main() {
	kv := map[string]StudentX{"menglu": {Age: 21}}
	a := kv["menglu"]
	a.Age = 22
	kv["menglu"] = a
	s := []StudentX{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kv, s)
}
