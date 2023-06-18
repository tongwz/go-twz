package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		tmpStu := stu
		m[stu.Name] = &tmpStu
	}

	fmt.Printf("我们m的值是： %#v \n", m)
	fmt.Printf("我们m的值是： %#v \n", m["wang"])
	fmt.Printf("我们m的值是： %#v \n", m["li"])
	fmt.Printf("我们m的值是： %#v \n", m["zhou"])
}

func main() {
	pase_student()
}
