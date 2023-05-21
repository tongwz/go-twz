package main

import (
	"fmt"
	model2 "projectTwo/model"
)

func main() {
	model := new(model2.Model)

	model.MyModel()

	var name string

	fmt.Scan(&name)
}
