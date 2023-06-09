package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhouJieLun"}}
	n := m["people"]
	n.name = "wuYanZu"
	m["people"] = n
	fmt.Println(n)
	fmt.Println(m)
}
