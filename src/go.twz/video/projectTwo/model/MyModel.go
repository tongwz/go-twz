package model

import "fmt"

type Model struct {
}

func (m *Model) MyModel() {
	fmt.Println("我是MyModel方法")
}
