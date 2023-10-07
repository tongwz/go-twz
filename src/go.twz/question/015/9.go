package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	// fmt.Println(m.Len())
	m.LoadOrStore("b", 2)
	m.LoadOrStore("c", 3)
	m.LoadOrStore("c", 2)

	var count int
	m.Range(func(key, value any) bool {
		count++
		return true
	})
	fmt.Println(count)
}
